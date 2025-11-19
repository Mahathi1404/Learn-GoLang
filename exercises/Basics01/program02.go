/*Sum two numbers passed by arguments*/
package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	nums:= os.Args[1:]
	var sum int
	if len(nums) < 2 {
		fmt.Println("Please provide two numbers")
		return
	}

	numA,err := strconv.Atoi(nums[0])
	if err != nil {
		fmt.Println("Invalid number:", nums[0],err)
		os.Exit(2) // terminates the program with exit code 2
	}

	numB,err := strconv.Atoi(nums[1])
	if err != nil {
		fmt.Println("Invalid number:", nums[1],err)
		os.Exit(2)
	}

	sum = numA + numB
	fmt.Printf("Sum of %d and %d is %d\n", numA,numB, sum)
}

// To check the exit status, run the program from terminal and then run 'echo $?' to see the last exit code

// Example usage:
// go run program02.go 5 10
// Output: Sum of 5 and 10 is 15
// echo $?
// Output: 0

// go run program02.go 5 abc
// Output: Invalid number: abc strconv.Atoi: parsing "abc": invalid syntax
// echo $?
// Output: 2

// echo $?, here $ indicates the starting of variable in shell, ? is the special variable that holds the exit status of the last executed command.

// Exit code 0 indicates successful execution
// Exit code 1 indicates general error
// Exit code 2 indicates syntax error in command arguments