# Structs, Methods and interfaces

## 1. Structs

A `struct` is a group of named fields of different types under a single logical unit. It allows to define custom data type.

## 1.1 Initialization
```go
type User struct {
    Name string
    id int
    Email string
}

var u User //Fields are set to zero values

v := User{"Alice", 34, "alice@mail.com"}

t := User{Name:"Bob", id:44, Email:"bob@hp.com"}

uPtr := new(User) //Allocates memory for the struct and returns a pointer (*User) to it

uPtr2 := &User{id: 2, Name: "Bob"}
```

## 1.2 Accessing Fields

```go
//access
fmt.Println(u.Name)

//Modify
u.id = 0

//accessing fields via pointer
uPtr := &u
uPtr.id = 98
```

## 1.3 Exporting and Visibility

Structs follow Go's standard visibility rules:

* Exported structs: Structure name must start with a capital letter.
* Exported Fields: Individual fields(Name, Email) must also start with a capital letter to be accessed or set from outside of package

A struct with an exported name but non-exported fields will be accessible, but its fields cannot be directly modified or read by an external package.

## 1.4 Anonymous structs

It is a feature that allows to define a structure and instantiate it in a single line, without giving it a formal type name.

They are essential for quick, temporary data models, but they come with a significant limitation: they cannot be reused easily.

```go
// a struct type with two fields (Name and Age) 
// and creating a variable 'person' of that type, all in one line.
person := struct {
    Name string
    Age  int
}{
    Name: "Alex", // Initialization of the fields
    Age:  30,
}

// Accessing fields is done normally:
fmt.Println(person.Name) // Output: Alex
```

Anonymous functions are used in situations where we need a qucik data container that will be used within a single function or for single transaction.
* Temporary data transfer(DTOs)
The most common use is creating a quick, structured data blob to pass to another library or to hold immediate results.
```go
func createLogEntry(action string, userID int) {
    // Define an anonymous struct for the log payload
    logPayload := struct {
        Time   time.Time
        Action string
        UserID int
    }{
        Time:   time.Now(),
        Action: action,
        UserID: userID,
    }

    // Pass the payload to a logging function that expects an interface{}
    logToService(logPayload)
}
```

* JSON marshalling/Unmarshalling : If you need to quickly parse a specific part of a JSON response without creating a persistent, named struct, an anonymous struct with the required fields and struct tags is perfect.
