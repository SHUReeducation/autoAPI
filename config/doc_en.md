# Config module design document

This document describes the basic idea of how the config is build up in autoAPI.

## What is config

See [our wiki](https://github.com/SHUReeducation/autoAPI/wiki/Config).

## Definitions

- layer: A certain kind of way to define some part of the config, eg. `config layer`, `CommandLineFlags layer`.

- module: A certain `struct` which is part of the config, eg. `Table`, `Database`

## Why it is hard to build a config

As you can see from the wiki, the config has 6 layers, and they can override config from each other.

To make things worse, some config can only be provided in several layers or one certain layer, 
eg. it is unreasonable to provide an SQL statement from the command line. 
And for some layers and/or modules,
whether it should work or not depends on another layer's and/or modules data.

To make things even worse, Go's package import system is really a piece of shit.
We cannot have cycles in the dependency graph. Which make it difficult to organize our packages.

So we need to find out how can we overcome these problems.

## Current Design

Current design is based on several rules:

- For each module and config source, we need to have functions named like `FromXXX`,
  eg. `FromYaml`, `FromFlags` to get values from a certain layer, and a default value named `Default`.
  
  The reason why we don't use an interface to normalize this is 
  - Interfaces in Golang doesn't require to be implemented explicitly. 
    We cannot use `impl XXX` to force some struct to implement some interface.
    And know the struct has to change its implementation if the interface is changed.
  - For some layers and/or modules, whether it should work or not depends on another layer's and/or modules data.
    So the interface is hard to give a good abstraction over these.
  - We do want to use concrete types instead of abstracted interfaces, 
    because use abstracted interfaces doesn't have any obvious benefits here.

- For each module, there should be a `MergeWith` function to merge two configs from different layers.
  General speaking, this function should looks like:
  
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

- The top-level module should join all configs in this order:
  ```go
  result := fromCommand
  result.MergeWith(fromEnv)
  result.MergeWith(fromconfig)
  result.MergeWith(fromSql)
  result.MergeWith(fromDatabase)
  result.MergeWith(Default)
  ```

- For each module, there should be a `Validate` function to check whether the module is valid.
  
  `Validate` will be called after all merges are completed.
  
- There occurse loop in dependency graph of modules:
  
  Instead of let the lower level of modules take/return instance of top-level module, 
  let it take/return top-level module's fields.