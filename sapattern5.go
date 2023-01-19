package main

import "fmt"

func main() {
	var i, j, k, rows int
	fmt.Println("Enter the number of rows: ")
	fmt.Scanln(&rows)
	fmt.Println("Mirrored Half Diamond Star Pattern")
	for i = 1; i <= rows; i++ {
		for j = rows - i; j > 0; j-- {
			fmt.Print(" ")
		}
		for k = 1; k <= i; k++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
	for i = rows - 1; i > 0; i-- {
		for j = rows - i; j > 0; j-- {
			fmt.Print(" ")
		}
		for k = 1; k <= i; k++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
