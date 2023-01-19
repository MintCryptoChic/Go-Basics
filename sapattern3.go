package main

import "fmt"

func main() {
	var i, j, k, l, rows int
	fmt.Println("Enter the number of rows: ")
	fmt.Scanln(&rows)
	fmt.Println("Hollow Diamond inside a Square Star Pattern")
	for i = 1; i <= rows; i++ {
		for j = i; j <= rows; j++ {
			fmt.Print("*")
		}
		for k = 1; k <= 2*i-2; k++ {
			fmt.Printf(" ")
		}
		for l = i; l <= rows; l++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
	for i = 1; i <= rows; i++ {
		for j = 1; j <= i; j++ {
			fmt.Print("*")
		}
		for k = 2*i - 2; k < 2*rows-2; k++ {
			fmt.Printf(" ")
		}
		for l = 1; l <= i; l++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
