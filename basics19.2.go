package main

import "fmt"

func main() {
	var x, y int
	fmt.Println("Enter the minimum number: ")
	fmt.Scanln(&x)
	fmt.Println("Enter the maximium number: ")
	fmt.Scanln(&y)
	fmt.Println("The natural numbers are: ")
	for x <= y {
		fmt.Println(x)
		x++
	}
}
