package main

import "fmt"

func main() {
	var num, prod int
	fmt.Println("Enter the number: ")
	fmt.Scanln(&num)
	fmt.Println("The product is: ")
	for prod = 1; num > 0; num = num / 10 {
		prod = prod * (num % 10)
	}
	fmt.Println(prod)
}
