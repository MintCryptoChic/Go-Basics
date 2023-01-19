package main

import "fmt"

func main() {
	var a, b, temp int
	fmt.Println("Enter the first number: ")
	fmt.Scanln(&a)
	fmt.Println("Enter the second number: ")
	fmt.Scanln(&b)

	temp = a
	a = b
	b = temp
	fmt.Println("The of first number is: ", a)
	fmt.Println("The of second number is: ", b)
}
