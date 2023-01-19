package main

import (
	"basics/yummy"
	"fmt"
)

func main() {
	var size int

	fmt.Print("Enter the array size: ")
	fmt.Scan(&size)
	fmt.Println("Enter the first array: ")
	arr1 := yummy.InputArray(size)
	fmt.Println("Enter the second array: ")
	arr2 := yummy.InputArray(size)
	add := make([]int, size)

	// fmt.Print("Enter the first array items: ")
	// for i := 0; i < size; i++ {
	// 	fmt.Scan(&arr1[i])
	// }

	// fmt.Print("Enter the second array items: ")
	// for i := 0; i < size; i++ {
	// 	fmt.Scan(&arr2[i])
	// }

	fmt.Print("Arrays after addition: ")
	for i := 0; i < size; i++ {
		add[i] = arr1[i] + arr2[i]
		fmt.Print(add[i], " ")
	}
}
