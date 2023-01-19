package main

import "fmt"

func main() {
	var num int
	fmt.Println("Enter the number: ")
	fmt.Scanln(&num)
	if num >= 0 {
		if num > 0 {
			fmt.Println("It is a positive number")
		} else {
			fmt.Println("The number is zero")
		}
	} else {
		fmt.Println("It is a negative number")
	}
}
