package main

import (
	"basics/yummy"
	"fmt"
)

func main() {
	var size int
	fmt.Println("Enter the size of the array: ")
	fmt.Scanln(&size)
	fmt.Println("Enter the first array: ")
	arr1 := yummy.InputArray(size)
	fmt.Println("Enter the second array: ")
	arr2 := yummy.InputArray(size)
	fmt.Println("Add\tSub\tMul\tDiv\tMod")
	for i := 0; i < size; i++ {
		fmt.Print(arr1[i]+arr2[i], "\t")
		fmt.Print(arr1[i]-arr2[i], "\t")
		fmt.Print(arr1[i]*arr2[i], "\t")
		fmt.Print(arr1[i]/arr2[i], "\t")
		fmt.Print(arr1[i]%arr2[i], "\t")
		fmt.Println()
	}
}
