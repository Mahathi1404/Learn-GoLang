/* enums*/

package main

import (
	"fmt"
	"reflect"
)

type Day uint8

const (
	Sunday Day = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

func main() {
	var today Day = Wednesday
	fmt.Println("Value of today is: ", today)
	fmt.Println("Type of today is: ", reflect.TypeOf(today))
}