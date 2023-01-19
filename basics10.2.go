package main

import "fmt"

func main() {
	var x, i int
	fmt.Println("Enter the number until which you want to print the even numbers: ")
	fmt.Scanln(&x)
	fmt.Println("The even numbers are: ")
	for i = 2; i <= x; i = i + 2 {
		fmt.Println(i)
	}
}
