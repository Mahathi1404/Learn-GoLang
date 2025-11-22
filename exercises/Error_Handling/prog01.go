/* exapmle program for error tpe, defer, panic and recover functions */

package main

import "fmt"

func riskyCode() {
	defer func() {
		r:= recover()

		if r!= nil {
			fmt.Println("Its okay, recovered from ", r)
		}
	}()

	panic("Something bad happened!")

	fmt.Println("End of risky code")
}

func main() {
	defer fmt.Println("Main function first defer")

	riskyCode()
	
	fmt.Println("End of main function")
}