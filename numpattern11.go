package main

import "fmt"

func main() {
	var i, j, rows int
	fmt.Println("Enter the number of rows: ")
	fmt.Scanln(&rows)
	fmt.Println("Right Angled Triangle with 1's and 0's")
	for i = 0; i < rows; i++ {
		for j = 0; j <= i; j++ {
			if j%2 == 0 {
				fmt.Print("1 ")
			} else {
				fmt.Print("0 ")
			}
		}
		fmt.Println()
	}
}
