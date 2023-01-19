package main

import "fmt"

func main() {
	var i, j, k, rows int
	fmt.Println("Enter the number of rows: ")
	fmt.Scanln(&rows)
	fmt.Println("Inverted Pyramid Star Pattern")
	for i = rows; i > 0; i-- {
		for j = 0; j <= rows-i; j++ {
			fmt.Print(" ")
		}
		for k = 0; k < (2*i - 1); k++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
