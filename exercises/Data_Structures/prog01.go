/*Write a function that takes an array of 5 integers and 
returns a new array with the elements in reverse order, gives its sum and maximum element*/

package main

import "fmt"

func revArray_sumElements(ip_arr [5]int) ([5]int, int, int){
	var rev_arr[5] int

	j:=0
	sum:=0
	maxElem:=ip_arr[0]
	for i:=len(ip_arr)-1;i>=0;i--{
		rev_arr[j]=ip_arr[i]
		sum+=ip_arr[i]
		j++
		if ip_arr[i]>maxElem{
			maxElem=ip_arr[i]
		}
	}
	return rev_arr, sum, maxElem
}

func main(){
	var arr[5] int =[5]int{4,5,6,7,8}
	fmt.Println("Original Array:", arr)
	
	result,s,m := revArray_sumElements(arr)
	fmt.Println("Reversed Array:", result, "Its Sum:", s, "Maximum Element:", m)
}