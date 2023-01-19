package main

import "fmt"

func main() {
	var i, j, rows int
	fmt.Println("Enter the number of rows: ")
	fmt.Scanln(&rows)
	fmt.Println("Right Angled Triangle Pattern")
	for i = 0; i < rows; i++ {
		for j = 0; j <= i; j++ {
			fmt.Print("1 ")
		}
		fmt.Println()
	}
}
