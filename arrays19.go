package main

import (
	"basics/yummy"
	"fmt"
)

func main() {
	var size, flag, ser int
	fmt.Println("Enter the size of the array: ")
	fmt.Scanln(&size)
	fmt.Println("Enter the elements of the array: ")
	arr1 := yummy.InputArray(size)
	fmt.Println("Enter the element to be searched: ")
	fmt.Scanln(&ser)
	for i := range arr1 {
		if ser == arr1[i] {
			flag++
			fmt.Println("The element was found at position: ", i)
			break
		} else {
			fmt.Println("Could not find the element in the array")
			break
		}
	}
}
