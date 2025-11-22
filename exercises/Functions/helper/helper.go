package help

import "fmt"

var HelperVar = getHelperVar()

func getHelperVar() string {
	fmt.Println("getHelperVar called")
	return "I am a helper variable"
}

func init() {
	fmt.Println("init function called from helper.go")
}

func init() {
	fmt.Println("init2 function called from helper.go")
}