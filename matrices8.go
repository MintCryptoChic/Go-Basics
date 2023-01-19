package main

import (
	"basics/yummy"
	"fmt"
)

func main() {
	var rows, columns int
	fmt.Println("Enter rows: ")
	fmt.Scanln(&rows)
	fmt.Println("Enter columns: ")
	fmt.Scanln(&columns)
	fmt.Println("Enter the matrix elements: ")
	matr1 := yummy.InputMatrix(rows, columns)
	fmt.Println("Upper Triangle: ")
	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			if i >= j {
				fmt.Print(matr1[i][j], "\t")
			} else {
				fmt.Print("0", "\t")
			}
		}
		fmt.Println()
	}
}
