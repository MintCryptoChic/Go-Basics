package main

import "fmt"

func main() {
	var num int
	fmt.Println("Enter the number: ")
	fmt.Scanln(&num)
	if num > 0 {
		fmt.Println("It is a positive number")
	} else if num < 0 {
		fmt.Println("It is a negative number")
	} else {
		fmt.Println("The number is zero")
	}
}
