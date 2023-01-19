package main

import "fmt"

func product(x int) int {
	var prod int
	for prod = 1; x > 0; x = x / 10 {
		prod = prod * (x % 10)
	}
	return prod
}

func main() {
	var num int
	fmt.Println("Enter the number: ")
	fmt.Scanln(&num)
	fmt.Println("The product is: ", product(num))

}
