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

### Functions

About Functions

- Most basic building block of Go programs
- Allows functionality to be isolated, which make programs easier to:
  - Test, debug, extend, modify, read, write, document
- Functions are simple: they take data as input and return data as output
  - Input and output data is optional

TODO creating functions

- Functions encapsulate program functionality which leads to more maintainable code
- Functions have parameters which define the input data
- Functions are used by calling the function and supplying arguments
- Functions can return multiple values
  - An underscore can be used to ignore a return value

### Operators

- Operators are used to perform calculations and comparisons
  - Less than `<`
- They work on operands
- Logic
  - `&&` And
  - `||` Or
  - `!` Not
- Logic operators can ensure multiple criteria are met
- Always returns a boolean
  - `authorized` && `authenticated`
  - `red` || `blue`
  - `!heavy`

### Flow Control

- Program code executes line-by-line
- Flow control is a way to interrupt this process
  - Different lines of code can be executed based on conditions
- Conditions are programmer-defined and can be as many as needed

```go
if temperature("freezer") > 0 {
  // display alert
}

if i := 5; i < 10 {}

if rank := getUserRank(); rank == "admin" {

} else if rank == "manager" {

} else {

}
```

Early Return

- Use "early return" to stop processing within a function as soon as possible
  - Max performance
  - All data needed for the function is ready after error checking
    - from "Clean Code"
