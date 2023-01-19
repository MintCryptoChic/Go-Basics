package main

import "fmt"

func main() {
	var x int
	fmt.Println("Enter the number until which you want to print even numbers: ")
	fmt.Scanln(&x)
	fmt.Println("The even numbers are: ")
	for i := 1; i <= x; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}
