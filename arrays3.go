package main

import (
	"fmt"
)

func main() {
	var size int
	fmt.Println("Enter the size of the arrays: ")
	fmt.Scanln(&size)
	arr1 := make([]int, size)
	fmt.Println("Enter the first array: ")
	for i := range arr1 {
		fmt.Scan(&arr1[i])
	}
	// arr1 := yummy.InputArray(size)
	arr2 := make([]int, size)

	fmt.Println("Enter the second array: ")
	for i := range arr2 {
		fmt.Scan(&arr2[i])
	}
	// arr2 := yummy.InputArray(size)
	arrm := make([]int, size)
	fmt.Print("Multiplication result: ")
	for i := 0; i < size; i++ {
		arrm[i] = arr1[i] * arr2[i]
		fmt.Print(arrm[i], "\t")
	}
}
