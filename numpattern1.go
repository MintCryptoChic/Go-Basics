package main

import "fmt"

func main() {
	var rows, columns, i, j int
	fmt.Println("Enter the number of rows: ")
	fmt.Scanln(&rows)
	fmt.Println("Enter the number of columns: ")
	fmt.Scanln(&columns)
	fmt.Println("Box Number Pattern of 1's and 0's")
	for i = 1; i <= rows; i++ {
		for j = 1; j <= columns; j++ {
			if i == 1 || i == rows || j == 1 || j == columns {
				fmt.Print("1 ")
			} else {
				fmt.Print("0 ")
			}
		}
		fmt.Println()
	}
}
