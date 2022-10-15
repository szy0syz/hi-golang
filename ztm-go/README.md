# ZTM - Go Programming

> 绝佳的练习英文的素材，边打边读。

## Introduction to Go

### Go CLI Tool

- Go toolchain provides the go command line utility
- This tool is used to:
  - Update dependencies
  - Build & test projects
  - Manage artifacts
- Format source code

> Everyday Go Commands 😂

- `build`: builds the projects & emits an executable binary
  - `build -race`: checks for concurrency problems
- `run`: runs the projects directly; no output executable
- `mod`: manages modules & dependencies
  - `mod tidy`: updates dependencies
- `test`: runs the project's test suite
- `fmt`: formats all source files (usually automated with IDE)

## Go Programming Fundamentals

### Variables

- Variables provide a way to store & access data in your application
  - Data within can be anything/vary (variable)
  - Alias to data in memory
  - Storing data to a variable is called assignment
- Variables have multiple components:
  - Name
  - Data (or lack thereof)
  - Type
  
```go
// Single Creation
var aa = 3

var aa int = 3

var aa int
aa = 3

// Compound Creation
var a, b, c = 1, 2, "sample"

var (
  a int = 1
  b int = 2
  c = "sample"
)

// Create & Assign
aa := 3

a, b := 1, "sample"
```

Defaults

- Variables that are declared but not assigned will automatically have a default value
  - `var name string`
  - String default: `""`
  - Number default: `0`
  - Other default: `nil`

Naming

- Go variable naming convention is camelCase
- Use names appropriate for the data

Constants

- Constant can be created using the const keyword
- Useful when declaring some value that needs to be utilized throughout some or all of the program

```go
const MaxSpeed = 30
const MinPurchasePrice = 7.50
const AppAuthor = "Bob"
```

Recap

- Variables are a way to access memory locations using an alias
- Multiple ways to create variables:
  - Single, compound, block, create & assign
- Variables can be assigned to other variables
- Variables names can only be used once per scope
- Variables declared, but not assigned to, will have a default value
- "Comma ok" idiom allows you to reuse the second variable
