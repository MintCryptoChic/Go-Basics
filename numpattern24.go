package main

import "fmt"

func main() {
	var i, j, rows int
	fmt.Println("Enter the number of rows: ")
	fmt.Scanln(&rows)
	fmt.Println("Right Arrow Number Pattern")
	for i = 1; i <= rows; i++ {
		for j = 1; j <= rows; j++ {
			if j < i {
				fmt.Print(" ")
			} else {
				fmt.Printf("%d", j)
			}
		}
		fmt.Println()
	}
	for i = 1; i < rows; i++ {
		for j = 1; j <= rows; j++ {
			if j < rows-i {
				fmt.Print(" ")
			} else {
				fmt.Printf("%d", j)
			}
		}
		fmt.Println()
	}
}
