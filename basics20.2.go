package main

import "fmt"

func main() {
	var num int
	fmt.Println("Enter the number for which you want to print natural numbers: ")
	fmt.Scanln(&num)
	fmt.Println("The natural numbers in reverse are: ")
	for num > 0 {
		fmt.Println(num)
		num--
	}
}
