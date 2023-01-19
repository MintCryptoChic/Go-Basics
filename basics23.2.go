package main

import "fmt"

func main() {
	var n int
	fmt.Println("Enter the number until which you want to print odd numbers: ")
	fmt.Scanln(&n)
	fmt.Println("The odd numbers are: ")
	for i := 1; i <= n; i = i + 2 {
		fmt.Println(i)
	}
}
