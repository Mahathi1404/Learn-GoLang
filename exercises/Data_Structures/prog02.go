/*Write a function that takes a slice of integers and returns a 
new slice containing only the even numbers from the original slice.*/

package main

import "fmt"

func getEvenElements(list []int) ([]int) {
	var evenEleArr []int
	for _,val := range list {
		if val%2==0 {
			evenEleArr = append(evenEleArr, val)
		}
	}
	return evenEleArr
}

func main() {
	numArr := []int{10, 15, 22, 33, 42, 55, 60}
	evenElements := getEvenElements(numArr)
	fmt.Println("Even elements in the array:", evenElements)
}