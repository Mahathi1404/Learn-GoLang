# Defer, Panic, Recover and Init function.

# Defer
The `defer` statement is used to postpone the function execution until its enclosed function finishes execution.

The behavior of `defer` with function parameters and closures relies on a single, crucial rule: The arguments to a deferred function are evaluated immediately, but the function body runs later

1. `defer` with Function Parameters (Arguments Evaluated Immediately)
When the Go runtime encounters a `defer` statement, it does not wait to calculate the arguments being passed to the `deferred` function. It calculates those argument values right then and there, saves those values, and then schedules the function call.

Example: Immediate Evaluation
In the example below, we might expect the defer to print x's value as 20, but because the argument is evaluated when defer is executed (when x is 10), it prints 10.

```go
func deferExample(){
    x:=10
    defer fmt.Println("deferred value of x is ",x)

    x=20
    fmt.Println("value of x (with normal execution)",x)
}

//output
//value of x (with normal execution) 20
//value of x is 10
```
2. `defer` with Closures (Closure Body Runs Later)
When you defer a closure, the rules change slightly because it is not deferring a value, it is deferring an entire function body.

A. The Closure Body Runs Later
When you defer a closure, the function body is what gets scheduled to run at the end of the surrounding function.

B. Access to the Latest Value
If the closure body accesses a captured variable (a variable defined in the outer scope), it accesses the latest value of that variable at the time the closure finally executes (just before the function returns).

Example: Latest Evaluation
In this example, we defer a closure. The closure itself doesn't evaluate x until the end of `latestEvaluation`.

```go
func latestEvaluation() {
    x := 10
    
    // 1. defer is encountered.
    // 2. The *function body* (the closure) is scheduled.
    // 3. This closure CAPTURES the variable x (its memory address/reference).
    defer func() {
        // This line runs at the end, accessing the latest value of x
        fmt.Printf("Deferred value of x: %d\n", x)
    }()

    x = 20 // The captured variable x is updated.
    
    fmt.Printf("Current value of x: %d\n", x)
}

// Output:
// Current value of x: 20
// Deferred value of x: 20 // The closure accessed the updated value
```

- LIFO Order: Deferred calls are executed in Last-In, First-Out (LIFO) order. The deferred call you write last is executed first.

- The common uses of `defer` include releasing resources and cleanup.

# Panic

A `panic` is a go tool that is used for dealing with unrecoverable, fatal error in the program. It stops the execution flow, executes deffered functions and returns control to the calling function.
If the panic reaches the top of the program (the main function) and is not stopped, the program crashes and prints a full stack trace.

Example:

```go
func panicExample() {
    defer fmt.Println("close panicExample()")

    fmt.Println("Hello from panicExample() function")

    //call panic
    panic("Panic here in example function")
}

func main() {
    defer fmt.Println("from main: Hello")

    panicExample()

    fmt.Println("Did we reach here")
}

//output
/*
Hello from panicExample() function
close panicExample()
from main: Hello
Panic here in example function
*/
```

In the above example, the last print statement in main is never reached.

# Recover

Recover is a process in Go to regain the control of the program when it hits the panic function.

`recover()` can only be used in deferred function within anonymous/closures function.

#### The Standard `recover` Pattern

```go
func startWorker() {
    // 1. Defer an anonymous function. This function runs when startWorker exits.
    defer func() {
        // 2. Call recover() and check if it returned a panic value (r != nil)
        if r := recover(); r != nil {
            // 3. A panic occurred! We handle it here instead of crashing.
            fmt.Printf("CATCH: Successfully recovered from fatal error: %v\n", r)
        }
    }() // Must include the '()' to execute the anonymous function immediately

    // This is the risky function that might call panic()
    riskyFunction() 
    
    // This line runs ONLY if riskyFunction() finished normally OR 
    // if a panic occurred but was successfully caught by the recover block.
}
```

Calling panic(i) executes the deferred function where the recover is different from nil. The returned value is the parameter of the panic function.

Note: If you try to call `recover()` directly as an argument to a deferred function, it fails because of the rule of Immediate Evaluation.

```go
func failsToRecover() {
    // 1. defer is encountered.
    // 2. The argument to fmt.Println, which is recover(), is evaluated IMMEDIATELY.
    // 3. Since no panic is active yet, recover() returns nil.
    defer fmt.Println("Recovered value (will be nil):", recover())

    panic("This will crash the program!") // The panic starts later.
    
    // Output:
    // Recovered value (will be nil): <nil> 
    // panic: This will crash the program!
    // (Program crashes)
}
```

Why it Fails: The `recover()` call executes instantly and sees no panic, returning `nil`. The panic then happens, but it has no chance to run the `recover()` function again.

# Init Functions

Every go program should have main package and starts execution with `main` function.
But libraries will not have main package. If library needs some initial configuration before being called, Go defines `init` functions which are executed once per package.

Init function are invoked automatically by go runtimes. A package can have multiple init functions
But they cannot be invoked from any part of code.

Syntax:
It takes no arguments or return any values
```go
func init() {
    //...
}
```

When we import a package the Go runtime follows this order:
1. Initialize imported packages recursively.
2. Initialize and assign values to variables.
3. Execute init functions.