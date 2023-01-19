package main

import "fmt"

func main() {
	var num int
	fmt.Println("Enter the number: ")
	fmt.Scanln(&num)
	for num >= 10 {
		num = num / 10
	}
	fmt.Println("The first digit is: ", num)
}
