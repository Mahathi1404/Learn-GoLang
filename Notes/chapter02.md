# Data types, Variables, Control flow and Functions

## 1. Data Types

A computer can understand only 1's and 0's, and all programs consist of binary numbers. A data type is a way for the machine to interpret these binary numbers, defining what kind of data the binary numbers represent. Common data types include numbers, characters, and strings.

Go is a statically typed language, meaning the data type must be specified before using variables. However, Go uses "type inference" to determine the type of data in certain circumstances. If the wrong data type is used, the compiler throws an error, providing type-checking features.

### Primitive Data Types in Go

Go supports several primitive data types, including numeric types, `bool`, and `string`. The type indicated in the program is a convention used for understanding the program.

#### Signed Integer Types:
- `int8`: -128 to 127
- `int16`: -32,768 to 32,767
- `int`, `int32`: -2,147,483,648 to 2,147,483,647

#### Unsigned Integer Types:
- `uint8`: 0 to 255
- `byte`: 0 to 255 (alias for `uint8`)
- `uint16`: 0 to 65,535
- `uint`, `uint32`: 0 to 4,294,967,295
- `uintptr`: 0 to `<pointer size>` (e.g., 0 to 2^64-1 on 64-bit systems)

#### Floating Point and Complex Types:
- `float32`, `float64`: Floating-point numbers
- `complex64`: 32-bit real and 32-bit imaginary parts
- `complex128`: 64-bit real and 64-bit imaginary parts

#### Rune
- `rune` Unicode code point (alias of int32)

#### Boolean Type:
- `bool`: Represents `true` or `false`

#### Type Aliases:
Type aliases allow assigning a new name to an existing type:
```go
type userId int
type Direction byte
type speed float64
type velocity speed
```

#### Type Conversion:
You can convert between types explicitly:
```go
userId(5) // Converts the integer 5 to type userId
speed(16.5) // Converts the float literal 16.5 to type speed
```

---

## 2. Variables

Variables are used to store and access data in a program. They act as references to data stored in memory. Storing data in a variable is called "assignment." Variables have three components: name, type, and data.

### Declaring Variables

Examples:
```go
var num = 3
var num int = 3

var num int
num = 3

arg:= 45

var isPresent bool
isPresent = false

isAbsent:= true

var name string = "amy"

comp:= "HP"

var aRune rune = '@' //rune represents character in UTF-8, it takes 32bits
```

### Default Values:
- Numeric types: `0`
- `bool`: `false`
- `string`: `""` (empty string)

### Multiple Variable Declaration:
```go
var a, b, c = 1, 2, "sample"

var (
    a int = 1
    b int = 2
    c     = "sample" // Type inference
)
```

### Short Variable Declaration:
```go
a := 3
a, b := 1, "sample"
```

### Reassignment:
Variables can be reassigned:
```go
a := 1
a = 2
a = 3 // Reassignment
```

### Scope Rules:
Redeclaration in the same scope is not allowed:
```go
a := 1
var a = 2 // Error
```

However, partial redeclaration is allowed:
```go
a, b := 1, 2
c, b := 3, 4 // OK
```

### Naming Conventions:
- Unexported (local) variables: `camelCase`
- Exported (public) variables: `PascalCase`

---

### Constants

Constants are immutable values defined using the `const` keyword. They must be assigned a value at the time of declaration, and this value cannot be changed later.

Example:
```go
const MaxSpeed = 30
const MinPurchase = 40
```

Naming conventions:
- Exported constants: `PascalCase` (e.g., `MaxSpeed`)
- Unexported constants: `camelCase` (e.g., `minPurchase`)

### Iota

The `const` keyword is used to define immutable values. It is common to group constants together, especially when they represent related states or options. The `iota` keyword in Go simplifies the assignment of incremental values to constants.

#### Basic Usage of `iota`
The `iota` keyword automatically generates incrementing values for constants within a block.

Example:
```go
const (
    Online      = 0
    Offline     = 1
    Maintenance = 2
)
```

Using `iota`, the same can be written as:
```go
const (
    Online = iota
    Offline
    Maintenance
)
```

Here, `iota` starts at `0` and increments by `1` for each subsequent constant.

#### Customizing `iota` Values
You can modify the value of `iota` by applying arithmetic operations.

Example:
```go
const (
    Start = iota + 3 // Start = 3
    Next             // Next = 4
    Final            // Final = 5
)
```

#### Mixing `iota` with Explicit Values
You can mix `iota` with explicitly assigned values, but `iota` will continue incrementing from the last value.

Example:
```go
const (
    Online      = iota // Online = 0
    Offline     = 1    // Explicitly set to 1
    Maintenance = iota // Maintenance = 2
)
```

---
## 3. Runes and Strings

### Runes

Text data in Go uses UTF-8 encoding. Encoding is a way of representing thousands of different symbols using code pages (think of it as a grid). Code pages are tables that use the first few bytes of data to determine which page to use. Each symbol in a code page is called a "code point."

For example:

|   | 1       | 2       | 3   | 4   |
|---|---------|---------|-----|-----|
| 1 | a (1,1) | d (1,2) |     |     |
| 2 |         |         |     |     |
| 3 |         |         |     |     |
| 4 |         |         |     | * (4F,1) |

- The letter `a` is at code point `(1,1)`. To create `a`, we need 2 bytes: `11`.
- The symbol `*` is at code point `(4F,1)`. To create `*`, we need 3 bytes: `4F1`.

Here, we are dealing with individual bytes, not the letters themselves.

Text is represented as the `rune` type in Go (similar to `char` in C). A `rune` is an alias for the `int32` type. A `rune` can represent any symbol, such as emojis or numbers. A `rune` is always a number, and a proper formatter is needed to print the corresponding characters.

Example:

```
    Q
_ _ _ 51     // The letter Q is represented as 1 byte
4 3 2 1
```

### Strings

A string is a data type for storing multiple runes. At a basic level, a string is just an array of bytes and its length. When iterating over a string, the iteration happens over bytes.

#### Strings in Memory

Example:

```
   Q
_ _ _ 51
4 3 2 1

   A
_ _ _ 11
4 3 2 1

QA
51 11
```

### Examples

- **Runes**: `'a'`, `'b'`, `'$'` (for special symbols)
- **Strings**: `"amount is $43\n"`
- **Raw Literal Strings**: `` `Let's code in "Golang"` ``


## 4. Control Flow

Control flow structures allow you to manage the execution of code based on conditions or loops.

### If-Else:
```go
if condition {
    // true block
} else if condition {
    // another condition
} else {
    // false block
}
```

### Switch:
Go's `switch` is more powerful than in many other languages:
- **Automatic Break**: No need for the `break` keyword.
- **Fallthrough**: Use the `fallthrough` keyword to execute the next case.
- **Expression-less Switch**: Acts as a clean replacement for long `if-else` chains.

Example:
```go
switch {
case value > 100:
    // ...
case value > 5:
    // ...
default:
    // ...
}
```

### Loops:

#### For Loop:
```go
for i := 0; i < 10; i++ {
    // ...
}
```

#### While-Style Loop:
```go
for i < 10 {
    i++
}
```

#### Infinite Loop:
```go
for {
    // ...
}
```

#### for range

Iterating over collections (strings, slices, maps).	
```go
for index, item := range collection { ... }
```

#### Breaking and Continuing:
- `break`: Exits the loop immediately.
- `continue`: Skips the current iteration.

Example:
```go
for {
    if condition {
        break
    }
}
```

## 5. Functions

Functions are the basic building blocks of a program. They allow functionality to be isolated, making a program easier to test, debug, modify, read, and write.

### Function Declaration

A function in Go is declared using the `func` keyword. Both input parameters and return types are optional.

Example:
```go
func main() {
    // Entry point of the program
}

func name(param1 type, param2 type) returnType {
    // Function body
}
```

### Example: Adding Two Numbers
```go
func sum(a, b int) int {
    return a + b
}
```
Here:
- `a` and `b` are arguments provided by the caller.
- The function returns the sum of `a` and `b` as an `int`.

### Multiple Return Values

Go functions can return multiple values.

Example:
```go
func multiReturn() (int, int, int) {
    return 1, 2, 3
}
```
This function returns three integers.

### Variadics

Functions that receive undetermined number of arguments are called variadic functions.

Example:
```go
func sum(nums ...int) int {
    //compute
    for _, data := range(nums) {
    }
}
//Here ...int tells that list of intergers will be passed
```

#### Function literals

Function can receive other functions as arguments, or assigned to a variable. Hence it is called First class citizens
Go permits anonymous functions and these functions can be closures.

A function literal means a function structure without a name after func, mainly defining a function locally. When this function literal is used without giving it the name, it becomes anonymous function.

```go
// The literal: func() { ... }
func() {
    fmt.Println("I run right away and have no name.")
}() // The parentheses here execute the function immediately.
//You write it, run it once, and forget about it. It’s useful for quick, isolated tasks.
```

If anonymous function capture variables, it is a closure

Closures can be understood this way. Most variables are public which are present in main function or the package which can be accessed by anyone. Closures create a private, personal workspace for its data and only closure itself can access or change the data.

It does it via function that remembers and protects its environment(the data) even after the function that created it is gone.

A closure is a block of code plus a set of external variables (the state). Both parts are treated as a single object in memory.
