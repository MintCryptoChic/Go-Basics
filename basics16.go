package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Println("Enter the first number: ")
	fmt.Scanln(&a)
	fmt.Println("Enter the second number: ")
	fmt.Scanln(&b)
	fmt.Println("Enter the third number: ")
	fmt.Scanln(&c)
	if (a > b) && (a > c) {
		fmt.Println("a is the greatest, value: ", a)
	} else if b > c {
		fmt.Println("b is the greatest, value: ", b)
	} else {
		fmt.Println("c is the greatest, value: ", c)
	}
}
