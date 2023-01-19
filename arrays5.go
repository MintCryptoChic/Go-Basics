package main

import (
	"basics/yummy"
	"fmt"
)

func main() {
	var size, count int
	fmt.Println("Enter the size of the array: ")
	fmt.Scanln(&size)
	fmt.Println("Enter the array elemnets: ")
	arr1 := yummy.InputArray(size)
	for i := 0; i < size; i++ {
		for j := i + 1; j < size; j++ {
			if arr1[i] == arr1[j] {
				count++
				break
			}
		}
	}
	fmt.Println("The total number of duplicates in the array are: ", count)
}
