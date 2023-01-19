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
	arr2 := make([]int, size)
	fmt.Println("Original array: ", arr1)
	j := 0
	for i := size - 1; i >= 0; i-- {
		arr2[j] = arr1[i]
		j++
	}
	fmt.Println("Reveresed array: ", arr2)
}
