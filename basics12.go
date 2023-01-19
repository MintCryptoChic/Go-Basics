package main

import "fmt"

func main() {
	var num, factorial int
	factorial = 1
	fmt.Println("Enter the number: ")
	fmt.Scanln(&num)
	for i := 1; i <= num; i++ {
		factorial = factorial * i
	}
	fmt.Println("The factorial of the number is: ", factorial)
}
