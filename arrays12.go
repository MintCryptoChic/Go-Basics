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
	fmt.Println("A list of array elements at odd index numbers: ")
	for i := range arr1 {
		if i%2 != 0 {
			fmt.Println(arr1[i])
		}
	}
}
