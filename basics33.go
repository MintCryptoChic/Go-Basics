package main

import "fmt"

func main() {
	var a, b int
	fmt.Println("Enter the number to be reversed: ")
	fmt.Scanln(&a)
	for a > 0 {
		b = b*10 + a%10
		a = a / 10
	}
	fmt.Println("The reversed number is: ", b)
}
