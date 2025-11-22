# Error Handling

Unlike other programming languages like python go doesnot have try/catch block, it is based on type `error`.
In Go, an error isn't an exception that stops the program flow; it's just another return value that must be checked.

To create error message we use `error.New()`

Example:

```Go
//Creating basic error
func getElementFromIndex(id int, data string) (int, error) {
    if id >= len(data) {
        // This creates an error value with the message.
        return 0, errors.New(fmt.Sprintf("Invalid index %d",id))
    }
}
```