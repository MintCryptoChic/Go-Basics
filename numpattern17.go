package main

import "fmt"

func main() {
	var i, j, rows int
	fmt.Println("Enter the number of rows: ")
	fmt.Scanln(&rows)
	fmt.Println("Inverted Right Triangle Pattern")
	for i = rows; i > 0; i-- {
		for j = i; j > 0; j-- {
			fmt.Print(1, " ")
		}
		fmt.Println()
	}
}
