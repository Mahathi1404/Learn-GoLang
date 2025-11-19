# Pointers

## 1. Introduction

A pointer is a variable that stores memory address of another variable.

## 1.1 Intialization of pointers

`var ptr *int` : declares a pointer variable called `ptr` that points address of integer value
`var s *string` : declares a pointer variable called `s` that points address of string data

## 1.2 Address-of and Dereferencing

Address-of operator (`&`), gives memory address of a variable.
Example:
    ```go
    x := 42
    p := &x //`p` now holds address of `x`
    ```

Dereference (`*`) : Gives the value stored at specific memory address.
Example:
    ```go
    *p // prints 42
    ```
