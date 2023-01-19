package main

import "fmt"

func main() {
	var i, j, rows, columns int
	fmt.Println("Enter the number of rows: ")
	fmt.Scanln(&rows)
	fmt.Println("Enter the number of columns: ")
	fmt.Scanln(&columns)
	fmt.Println("Rectangle Pattern")
	for i = 0; i < rows; i++ {
		for j = 0; j < columns; j++ {
			fmt.Print(1, " ")
		}
		fmt.Println()
	}
}
