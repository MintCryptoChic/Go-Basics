package main

import (
	"basics/yummy"
	"fmt"
)

func main() {
	var size, evensum, oddsum int
	fmt.Println("Enter the size of the array: ")
	fmt.Scanln(&size)
	fmt.Println("Enter the elements of the array: ")
	arr1 := yummy.InputArray(size)
	for i := 0; i < size; i++ {
		if arr1[i]%2 == 0 {
			evensum++
		} else if arr1[i]%2 != 0 {
			oddsum++
		}
	}
	fmt.Println("Number of even numbers: ", evensum)
	fmt.Println("Number of odd numbers: ", oddsum)
}
