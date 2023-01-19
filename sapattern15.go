package main

import "fmt"

func main() {
	var i, j, k, rows int
	fmt.Println("Enter the number of rows: ")
	fmt.Scanln(&rows)
	fmt.Println("Left Pascals Star Triangle Pattern")
	for i = 1; i <= rows; i++ {
		for j = i; j < rows; j++ {
			fmt.Print(" ")
		}
		for k = 1; k <= i; k++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
	for i = rows - 1; i > 0; i-- {
		for j = i; j < rows; j++ {
			fmt.Print(" ")
		}
		for k = 1; k <= i; k++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
