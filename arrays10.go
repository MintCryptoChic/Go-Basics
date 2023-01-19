package main

import (
	"fmt"
)

func main() {
	var size int
	fmt.Println("Enter the size of the array: ")
	fmt.Scanln(&size)
	fmt.Println("Enter the elements of the array: ")
	arr1 := make([]int, size)
	for i := 0; i < size; i++ {
		fmt.Scanln(&arr1[i])
	}
	fmt.Println("The elements of the array with their index numbers are: ")
	for i := range arr1 {
		fmt.Println(arr1[i], "at index number", i)
	}
}
