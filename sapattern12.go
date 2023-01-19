package main

import "fmt"

func main() {
	var i, j, k, rows int
	fmt.Println("Enter the number of rows: ")
	fmt.Scanln(&rows)
	fmt.Println("Star Pyramid Pattern")
	for i = 1; i <= rows; i++ {
		for j = 1; j <= rows-i; j++ {
			fmt.Print(" ")
		}
		for k = 1; k <= 2*i-1; k++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
