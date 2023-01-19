package main

import (
	"basics/yummy"
	"fmt"
)

func main() {
	var size, pos, neg, z int
	fmt.Println("Enter the size of the array: ")
	fmt.Scanln(&size)
	fmt.Println("Enter the elements of the array: ")
	arr1 := yummy.InputArray(size)
	for i := 0; i < size; i++ {
		if arr1[i] < 0 {
			neg++
		} else if arr1[i] > 0 {
			pos++
		} else if arr1[i] == 0 {
			z++
		}
	}
	fmt.Println("Positive numbers: ", pos)
	fmt.Println("Negative numbers: ", neg)
	fmt.Println("Zeroes: ", z)
}
