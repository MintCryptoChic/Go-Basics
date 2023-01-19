package main

import "fmt"

func main() {
	var i, j, k, rows int
	fmt.Println("Enter the number of rows: ")
	fmt.Scanln(&rows)
	fmt.Println("Same Numbers in Square Rows and Columns Pattern")
	for i = 1; i <= rows; i++ {
		for j = i; j <= rows; j++ {
			fmt.Printf("%d ", j)
		}
		for k = 1; k < i; k++ {
			fmt.Printf("%d ", k)
		}
		fmt.Println()
	}

}
