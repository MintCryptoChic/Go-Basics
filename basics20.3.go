package main

import "fmt"

func main() {
	var num1, num2 int
	fmt.Println("Enter the minimum number for which you want to print natural numbers: ")
	fmt.Scanln(&num1)
	fmt.Println("Enter the maximum number for which you want to print natural numbers: ")
	fmt.Scanln(&num2)
	fmt.Println("The natural numbers in reverse are: ")
	for num2 >= num1 {
		fmt.Println(num2)
		num2--
	}
}
