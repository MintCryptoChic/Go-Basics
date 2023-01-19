package main

import (
	"basics/yummy"
	"fmt"
)

func main() {
	var size, p, n int
	fmt.Println("Enter the size of the array: ")
	fmt.Scanln(&size)
	fmt.Println("Enter the elements of the array: ")
	arr1 := yummy.InputArray(size)
	arrp := make([]int, size)
	arrn := make([]int, size)
	for i := range arr1 {
		if arr1[i] > 0 {
			arrp[p] = arr1[i]
			p++
		} else if arr1[i] < 0 {
			arrn[n] = arr1[i]
			n++
		}
	}
	fmt.Println("Array of positive values: ")
	for i := 0; i < p; i++ {
		fmt.Print(arrp[i], "\t")
	}
	fmt.Println()
	fmt.Println("Array of negative values: ")
	for i := 0; i < n; i++ {
		fmt.Print(arrp[i], "\t")
	}
}
