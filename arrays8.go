package main

import (
	"basics/yummy"
	"fmt"
)

func main() {
	var size, pos int
	fmt.Println("Enter the size of the array: ")
	fmt.Scanln(&size)
	fmt.Println("Enter the elements of the array: ")
	arr1 := yummy.InputArray(size)
	temp := arr1[0]
	for i := 0; i < size; i++ {
		if temp < arr1[i] {
			temp = arr1[i]
			pos = i
		}
	}
	fmt.Println("The largest number is: ", temp)
	fmt.Println("The index number of the largest number is: ", pos)

}
