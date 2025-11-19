/* Constant declaration*/

package main

import (
	"fmt"
	"reflect"
)

const (
	PI = 3.14
	AvagadroNum float32 = 6.022e23 
)

func main() {
	fmt.Println("Value of PI is: ",PI)
	fmt.Println("Type of PI is: ", reflect.TypeOf(PI))
	
	fmt.Println("Value of AvagadroNum is: ",AvagadroNum)
	fmt.Println("Type of AvagadroNum is: ", reflect.TypeOf(AvagadroNum))
}