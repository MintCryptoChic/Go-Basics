package main

import "fmt"

func fac(x int) int {
	fact := 1
	for i := 1; i <= x; i++ {
		fact = fact * i
	}
	return fact
}

func main() {
	var num int
	fmt.Println("Enter the number: ")
	fmt.Scanln(&num)
	fmt.Println("The factorial is: ", fac(num))
}
