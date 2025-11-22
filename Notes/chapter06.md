# Arrays, Slices and Maps

## 1. Arrays

An array is a linear data structure which stores similar type of data.
In Go, arrays are typed and their size is fixed that means size cannot be changed once it is created.

### 1.1 Different ways of array declarations

```go
var a[5] int //by default it will store 0
fmt.Println(a) //[0,0,0,0,0]

b := [5]int{0,1,2,3,4}
fmt.Println(b) //[0,1,2,3,4]

c := [5]int{1,2,3}
fmt.Println(c) //1,2,3,0,0
fmt.Println(len(c)) //5

//if you want compiler to count the values in the array and fix the size

d := [...]{"a","b","c"}
```

## 2. Slices

Slices are a way to handle dynamic ordered list of data.

Slice is a dynamically-sized view into underlying array. It is defined as descriptor for a contiguous segment of an underlying array and provides access to a numbered sequence elements from that array.

Slice is like a way of telling "The data you want starts at this exact spot in the memory and runs unbroken for this many items"

Slices are fast and flexible because they just change the descriptor to grow or shrink without having to move or copy the array everytime.

A slice descriptor is a small structure that contain three components:

* Pointer : A pointer to the first elements of the array
* Length(`len`) : The number of elements that are accessible
* Capacity(`cap`) : The number of elements in the underlying array, starting from the slice's pointer, that are available to the slice without reallocating a new underlying array.

### 2.1 Declaring and initializing slices

```go
s := []int{10,20,30}

//using make function
//make(data type, len,cap)
a := make([]int,5,10)

b := make([]string,3)

//Creating slice from an array or other slice

var arr[10] int;
arr = {1,2,3,4,5,6,7,8,9}

s1:= arr[3:5] //includes element from 3rd index upto 5th(excluding) len:2,cap:7(considers till end of array)
s2:= arr[1:] //creates slice from second element to last
s4:= arr[:] //creates a slice covering entire source
```

### 2.2 Append function

`append` adds new values to the end of a slice.

If the slice has enough capacity, the new element is added to the underlying array, and the slice's length is increased.

If the slice is full (length equals capacity), a new, larger underlying array is created, all existing elements are copied to the new array, and the new elements are added. The slice descriptor is updated to point to the new array.

```go
sl := []int{1,2} //len 2, cap 2

append(sl,3) //len :3 , cap:4 (since slice was full, cap doubles)

append(sl,4,5,6) //len:6,cap:8
```

### 2.3 Copying Slices

The `copy` function is used to copy elements from source slice to destination slice

Syntax: `copy(dest,src)`

```go
src := []int{2,3,4}
dst := make([]int,2)

copy(dst,src) //copies first two elements of src to dst
```

Note: Since slice is just view of underlying array, multiple slices can share the same underlying array. Modifying element in one slice will change in other slices as well.
Any element with a position equal to or greater than length cannot be accessed independently of the slice capacity

```go
// Underlying array: {10, 20, 30, 40}
s1 := []int{10, 20, 30, 40}

// s2 points to the same underlying array segment as s1
s2 := s1[1:3] // s2 is {20, 30}

s1[2] = 99 // Change the element at index 2 in s1 (which is the element '30')

// The change is reflected in s2 because they share the underlying data
fmt.Println(s2) // Output: [20 99]
```

### 2.4 Iterating array,slices

```go
for i:=0;i<len(arr);i++{
    //...
}

for idx,val := range arr{
    //prints idx and its data
}
```

## 3. Maps

Map is a unordered data structure that stores key value pair, where keys are unique entity.
Slices and maps cannot be used as map keys.

Reference Type: Like slices, maps are reference types. When a map is passed to a function or assign it to a new variable, it is passing a reference to the underlying data structure, not a copy of the data. Changes made through one reference affect all others.

### 3.1 Initializing maps

Syntax: `mp := make(map[keyType]valueType)`

```go
frutis := make(map[string]int)

statusCodes := map[int]string {
    200: "OK",
    404: "Not found",
    500: "Interval server Error"
}

/*Note : Declaring a map without using make results in a nil map. We cannot add elements to a nil map; it will cause a runtime panic.*/
var scores map[string]int //This initializes map to nil
//scores["Al"] = 95 //Causes runtime error
```

### 3.2 Map operations

1. Adding or updating elements

```go
p["Tokyo"] = 2
p["Tokyo"] = 3
```

2. Read elements

```go
val := p["tokyo"]
```

3. Checking for existence of element

```go
val, exists := p["tokyo"]
if exists {
    fmt.Println("Tokyo is present")
} else {
    fmt.Println("Not present")
}
```

4. Delete elements

```go
delete(p, "Tokyo")
```

### 3.3 Iterating

```go
for key,value := range statusCodes {
    fmt.Println(key,value)
}

//for iterating only keys
for code := range statusCodes {
    fmt.Println(code)
}
```