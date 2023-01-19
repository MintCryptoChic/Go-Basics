package main

import "fmt"

func main() {
	var i, j, rows int
	fmt.Println("Enter the number of rows: ")
	fmt.Scanln(&rows)
	fmt.Println("Right Triangle of Numbers in Reverse Pattern")
	for i = rows; i > 0; i-- {
		for j = rows; j >= i; j-- {
			fmt.Printf("%d ", j)
		}
		fmt.Println()
	}
}
