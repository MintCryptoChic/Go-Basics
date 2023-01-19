package main

import "fmt"

func main() {
	var i, j, rows int
	fmt.Println("Enter the number of rows/columns: ")
	fmt.Scanln(&rows)
	fmt.Println("Square Number Pattern")
	for i = 1; i <= rows; i++ {
		for j = 1; j <= rows; j++ {
			fmt.Printf("1 ")
		}
		fmt.Println()
	}
}
