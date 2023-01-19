package main

import (
	"basics/yummy"
	"fmt"
)

func main() {
	var size, l, s, p, q int
	fmt.Println("Enter the size of the array: ")
	fmt.Scanln(&size)
	fmt.Println("Enter the array elemnets: ")
	arr1 := yummy.InputArray(size)
	l = arr1[0]
	s = arr1[0]
	for i := 0; i < size; i++ {
		if arr1[i] > l {
			l = arr1[i]
			p = i
		}
		if arr1[i] < s {
			s = arr1[i]
			q = i
		}
	}
	fmt.Println("The largest number is: ", l, "at position: ", p)
	fmt.Println("The smallest number is: ", s, "at position: ", q)
}
