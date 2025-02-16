/* Program to understand function literals and closures in Go */
package main

import (
	"fmt"
)

func main() {
	// Function literal assigned to a variable
	add := func(a int, b int) int {
		return a + b
	}

	// Using the function literal
	result := add(3, 5)
	fmt.Println("Sum:", result)

	// Closure example
	counter := func() func() int {
		count := 0
		return func() int {
			count++
			return count
		}
	}()

	// Using the closure
	fmt.Println("Counter:", counter()) // Output: 1
	fmt.Println("Counter:", counter()) // Output: 2
	fmt.Println("Counter:", counter()) // Output: 3
}