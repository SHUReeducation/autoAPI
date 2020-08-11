# Config module design document

This document describes the basic idea of how the config is build up in autoAPI.

## What is config

See [our wiki](https://github.com/SHUReeducation/autoAPI/wiki/Config).

## Definitions

- layer: A certain kind of way to define some part of the config, eg. `configFile layer`, `CommandLineFlags layer`.

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

- Basically, we construct the config info of the highest level, and merge it with lower levels, level by level.
  
  ie.
  ```go
  currentConfigure := config.FromCommandLine(commandLine)
  currentConfigure.MergeWithEnv()
  currentConfigure.MergeWithconfig(commandLine.config)
  currentConfigure.MergeWithSQL(commandLine.sqlPath)
  currentConfigure.MergeWithDB()
  currentConfigure.MergeWithDefaultValue()
  ```
  
  The reason why we don't use an interface to normalize the merge behaviour of merging is 
  - Interfaces in Golang doesn't require to be implemented explicitly. 
    We cannot use `impl XXX` to force some struct to implement some interface.
    And know the struct has to change its implementation if the interface is changed.
  - For some layers and/or modules, whether it should work or not depends on another layer's and/or modules data.
    So the interface is hard to give a good abstraction over these.
  - We do want to use concrete types instead of abstracted interfaces, 
    because use abstracted interfaces doesn't have any obvious benefits here.

- For each module, there should be a `Validate` function to check whether the module is valid.
  
  `Validate` will be called after all merges are completed.
  
- There occurse loop in dependency graph of modules:
  
  Instead of let the lower level of modules take/return instance of top-level module, 
  let it take/return top-level module's fields.