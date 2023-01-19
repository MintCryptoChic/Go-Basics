package main

import (
	"basics/yummy"
	"fmt"
)

func main() {
	var size, sum int
	fmt.Println("Enter the size of the array: ")
	fmt.Scanln(&size)
	fmt.Println("Enter the array elemnets: ")
	arr1 := yummy.InputArray(size)
	for i := range arr1 {
		sum = sum + arr1[i]
	}
	avg := float64(sum / size)
	fmt.Println("The sum of the elements: ", sum)
	fmt.Println("The average of the elements: ", float64(avg))
}
