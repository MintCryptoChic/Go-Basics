package main

import "fmt"

func fdigit(x int) int {
	for x >= 10 {
		x = x / 10
	}
	return x
}

func main() {
	var num int
	fmt.Println("Enter the number: ")
	fmt.Scanln(&num)
	fmt.Println("The first digit is: ", fdigit(num))
}
