package main

import "fmt"

func main() {
	var i, j, rows int
	fmt.Println("Enter the number of rows: ")
	fmt.Scanln(&rows)
	fmt.Println("Inverted Right Triangle Pattern")
	for i = 1; i <= rows; i++ {
		for j = rows; j >= i; j-- {
			fmt.Printf("%d ", j)
		}
		fmt.Println()
	}
}
