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

现有设计基于以下几个规定：

- 对每组 (layer, module)，我们需要定义函数`FromXXX`,
  如 `FromYaml`， `FromFlags`来从某个特定 layer 构造module， 如果有默认值的话可以称之为 `Default`.
  
  我们不使用一个 `interface` 来规范这些的原因是：
  - 无法显式实现 Golang 中的接口，即我们不能通过 `impl XXX` 来强制某个结构体实现某个接口，并在接口改变时得到编译错误。
  - 对于部分 layers 和/或 modules, 其是否生效取决于其他 layer 和/或 modules 中的值，定死了接口传入这些额外参数就成为了一个难题。
  - 我们确确实实在这里使用了具体类型而非抽象的接口，在这里用接口也没得到什么好处。

- 对每个 module， 应该有一个 `MergeWith` 函数来将两个不同 layer 的配置合并起来。
  这个函数一般来说应该像这样：
  
  ```go
  func (t *T) MergeWith(other *T) {
    if other == nil {
      return
    }
    if t.Field1 == nil {
      t.Field1 = other.Field1
    }
    if t.Field2 == nil {
      t.Field2 = other.Field2
    }
    // ...
    t.SubModule1.MergeWith(other.SubModule1)
    t.SubModule2.MergeWith(other.SubModule2)
    // ...
  }
  ```

- 顶层 module 应该将其不同layer获得的配置全部合并起来
  ```go
  result := fromCommand
  result.MergeWith(fromEnv)
  result.MergeWith(fromConfigFile)
  result.MergeWith(fromSql)
  result.MergeWith(fromDatabase)
  result.MergeWith(Default)
  ```

- 每个 module 应该提供 `Validate` 函数来检查其最终合法性。
  
  Validate 将在所有 merge 都完成后调用。
  
- 如果两个 module 之间有构成循环引用的危险，解决方案：

  让底层的 module 不再接受或返回顶层 module 的实例，而接受/返回顶层 module 的字段