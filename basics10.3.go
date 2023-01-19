package main

import "fmt"

func main() {
	var x int
	fmt.Println("Enter the number until which you want to print the even numbers: ")
	fmt.Scanln(&x)
	fmt.Println("The even numbers are: ")
	i := 2
	for i <= x {
		fmt.Println(i)
		i = i + 2
	}
}
