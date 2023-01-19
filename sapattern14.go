package main

import "fmt"

func main() {
	var i, j, k, rows int
	fmt.Println("Enter the number of rows: ")
	fmt.Scanln(&rows)
	fmt.Println("Hollow Inverted Star Pyramid Pattern")
	for i = rows; i > 0; i-- {
		for j = 1; j <= rows-i; j++ {
			fmt.Print(" ")
		}
		for k = 1; k <= 2*i-1; k++ {
			if k == 1 || k == 2*i-1 || i == 1 || i == rows {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
