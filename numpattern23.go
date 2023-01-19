package main

import "fmt"

func main() {
	var i, j, rows int
	fmt.Println("Enter the number of rows: ")
	fmt.Scanln(&rows)
	fmt.Println("Right Pascals Number Triangle Pattern")
	for i = 1; i <= rows; i++ {
		for j = 1; j <= i; j++ {
			fmt.Printf("%d ", j)
		}
		fmt.Println()
	}
	for i = rows - 1; i > 0; i-- {
		for j = 1; j <= i; j++ {
			fmt.Printf("%d ", j)
		}
		fmt.Println()
	}
}
