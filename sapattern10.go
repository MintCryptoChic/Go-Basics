package main

import "fmt"

func main() {
	var i, j, rows int
	fmt.Println("Enter the number of rows: ")
	fmt.Scanln(&rows)
	fmt.Println("Hollow Inverted Right Triangle Star Pattern")
	for i = rows; i > 0; i-- {
		for j = i; j > 0; j-- {
			if j == 1 || i == 1 || i == rows || j == i {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
