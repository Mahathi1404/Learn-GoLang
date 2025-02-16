# 1. Introduction to Go

## 1.1 What is Golang?

Golang (or Go) is an open-source, statically typed, compiled programming language developed by Google.

**Main features:**
- **Speed and safety** Go is statically typed meaning types are checked at compile time
- **Readability:** Easy to learn and write programs
- **High performance:** Efficient execution

## 1.2 Design

Go was mainly designed with four principles:

- **Simplicity:** Minimal syntax, fewer keywords. Avoids features like classes and inheritance.
- **Efficiency:** Compiled to native machine code for fast execution. Includes a built-in garbage collector for automatic memory management.
- **Concurrency:** Built-in, easy-to-use concurrency support using goroutines and channels.
- **Tooling:** Many tools are available within the `go` command.

## 1.3 Go Usage

Go is most commonly used in backend infrastructure, such as:

- Cloud infrastructure (e.g., Kubernetes, Docker)
- Networking APIs
- Command-line interfaces (CLIs)

# 2. Basic structure of a Go program

## 2.1 The "Hello, World" program

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, world!")
}
```

## 2.2 Understanding the code

- `package main` — declares this file as part of the `main` package; a `main` package produces an executable.
- `import "fmt"` — imports the `fmt` package, which provides formatted I/O functions.
- `func main()` — the program entry point. Execution starts from `main` in the `main` package.
- `fmt.Println()` — prints the provided arguments followed by a newline.

## 2.3 How to run Go programs

Two common ways:

1. `go run program.go` — compiles and runs the program in one step (no binary kept).
2. `go build program.go` — compiles and produces an executable binary you can run separately (e.g., `./program` on Unix-like systems).

# 3. Code organization

Go projects are organized hierarchically using directories, where each directory typically represents a single package.

## 3.1 Packages

Packages is the primary way to organize and reuse code.
A package is a directory that contains one or more Go source files that share the same package name declaration.

The `package main`
Purpose: This special package is reserved for executable programs. It is the entry point of the application. If package is main then compiler will understand that it is executable file not just sharable library

`import` keyword brings external packages to go code

The `import "fmt"` statement brings in the fmt package

If importing from github, `import "github.com/gin-gonic/gin"`

Packages are used for modularity, encapsulation and compilation. Every .go file must belong to a package.

Encapsulation (public/private entities): Packages control visibility
Any identifier(functions, struct, variables) that starts with a capital letter is made public(can be exported) and visible outside directory

Identifier starting with lowercase letter cannot be exported i.e private, can only be used within the package.

Note: If building library then package name is recommended to be same as its folder name.

## 3.2 Modules

Modules are group of one or more packages. It defines dependencies and versioning.
It is defined by the file `go.mod`, present in the top-level directory.

## This is how a Go project look like.

my_project/  <-- module root (has go.mod)
├── go.mod
├── main.go          <-- package main
├── utils/           <-- PACKAGE DIRECTORY
│   └── helper.go    <-- package utils
└── models/          <-- PACKAGE DIRECTORY
    └── user.go      <-- package models

## 3.3 Project initialization

A Go project is initialized by creating a Go Module, which enables dependency management.

`mkdir my_project && cd my_project`
 
`go mod init myFirstProject/dir_name`
Here 'myFirstProject' is the logical name that is given by user

`go mod init github.com/yourusername/my_project`

The above command will create go.mod file inside the directory

Importing the packages:

> Alias the 'utils' package from the long module path to 'us':

import us "github.com/yourusername/my_project/utils" 
**Usage inside code:**
// Usage inside code:
us.HelperFunction()

`go get <package_name>` is used to add dependencies to your module (as of Go 1.17+, use `go install <package>@latest` to install executables globally).
`go mod tidy` removes unused dependencies, adds any missing ones, and ensures both `go.mod` and `go.sum` files are accurate.

`go install` installs the executable binary to `$GOPATH/bin` or `$GOBIN`. In module-aware mode (Go 1.17+), the recommended usage is `go install pkg@version` to install a specific version globally.