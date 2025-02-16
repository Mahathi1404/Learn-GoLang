/*Program to pass arguments to our program*/

package main

import (
	"fmt"
	"os"
)

func main() {
	args_w_prog := os.Args // return array containg arguments passed to the program
	args_wo_prog := os.Args[1:]

	fmt.Println("Aruguments with program ",args_w_prog)
	fmt.Println("Arguments without programe name ",args_wo_prog)
}