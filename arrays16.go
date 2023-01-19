package main

import (
	"basics/yummy"
	"fmt"
)

func main() {
	var size int
	fmt.Println("Enter the size of the array: ")
	fmt.Scanln(&size)
	fmt.Println("Enter the elements of the array: ")
	arr1 := yummy.InputArray(size)
	fmt.Println("A list of positive numbers in the array: ")
	for i := range arr1 {
		if arr1[i] > 0 {
			fmt.Println(arr1[i])
		}
	}
}
