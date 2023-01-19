package main

import "fmt"

func main() {
	var n int
	fmt.Println("Enter the number until which you want to print odd numbers: ")
	fmt.Scanln(&n)
	fmt.Println("The odd numbers are: ")
	for i := 0; i <= n; i++ {
		if i%2 != 0 {
			fmt.Println(i)
		}
	}
}
