package main

import (
	"basics/yummy"
	"fmt"
)

func main() {
	var size, sum int
	fmt.Println("Enter the size of the array: ")
	fmt.Scanln(&size)
	fmt.Println("Enter the elements of the array: ")
	arr1 := yummy.InputArray(size)
	for i := range arr1 {
		sum = sum + arr1[i]
	}
	fmt.Println("Sum of array items: ", sum)
}
