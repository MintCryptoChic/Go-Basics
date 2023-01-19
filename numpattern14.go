package main

import "fmt"

func main() {
	var i, j, rows int
	fmt.Println("Enter the number of rows: ")
	fmt.Scanln(&rows)
	fmt.Println("Right Triangle of Incremented Numbers Pattern")
	for i = 1; i <= rows; i++ {
		for j = i; j > 0; j-- {
			fmt.Printf("%d ", j)
		}
		fmt.Println()
	}
}
