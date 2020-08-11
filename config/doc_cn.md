# 配置设置模块设计文档

该文档介绍了 autoAPI 中整个配置设置的生成方式。

## 配置设置是什么

见 [我们的 wiki](https://github.com/SHUReeducation/autoAPI/wiki/%E9%85%8D%E7%BD%AE).

## 定义

- layer: 某种定义配置设置的方式，例如 `配置文件 layer`, `命令行 layer`.

- module: 某个特定的代表配置设置的某一块的 `struct`，例如 `Table`, `Database`

## 配置设置生成的难点

正如 wiki 所述，配置来源有 6 个 layer，且它们可以互相覆盖。

更糟糕的是，有些配置项只能在部分 layer 中提供，
例如让用户在命令行提供 sql 语句是不切实际的。
且有些 layer 是否生效取决于其他 layer 中的值。

更更糟糕的是 Golang 这傻◻语言不支持循环引用。

所以我们必须找到一个克服以上问题的方法。

## 现有设计

- 大体上，先构造出优先级最高的 `layer` 的总配置信息，再逐层向下构造并合并。
    
  即：
  ```go
  currentConfigure := config.FromCommandLine(commandLine)
  currentConfigure.MergeWithEnv()
  currentConfigure.MergeWithconfig(commandLine.config)
  currentConfigure.MergeWithSQL(commandLine.sqlPath)
  currentConfigure.MergeWithDB()
  currentConfigure.MergeWithDefaultValue()
  ```
  
  我们不使用 interface 规范每个层之间的 merge 行为的原因是：
    - Golang 的 interface 并没有强制性。我们不能用 `impl XXX` 来强制某个 `struct` 实现某个 `interface`.
    - 有些模块的有些层需要一些来自顶层模块的参数，所以这个 `interface` 抽象不出来。
    - 这里引入一个 `interface` 没有啥好处 —— 我们还是用的具体的类型。

- 每个 module 应该提供 `Validate` 函数来检查其最终合法性。
  
  Validate 将在所有 merge 都完成后调用。
  
- 如果两个 module 之间有构成循环引用的危险，解决方案：

  让底层的 module 不再接受或返回顶层 module 的实例，而接受/返回顶层 module 的字段
