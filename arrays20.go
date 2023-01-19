package main

import (
	"basics/yummy"
	"fmt"
)

func main() {
	var size, small, pos int
	fmt.Println("Enter the size of the array: ")
	fmt.Scanln(&size)
	fmt.Println("Enter the elements of the array: ")
	arr1 := yummy.InputArray(size)
	small = arr1[0]
	for i := range arr1 {
		if small > arr1[i] {
			small = arr1[i]
			pos = i
		}
	}
	fmt.Println("Smallest value in array: ", small, "at position: ", pos)
}
