package main

import (
	"basics/yummy"
	"fmt"
)

func main() {
	var rows, columns, sum int
	fmt.Println("Enter rows: ")
	fmt.Scanln(&rows)
	fmt.Println("Enter columns: ")
	fmt.Scanln(&columns)
	fmt.Println("Enter the matrix elements: ")
	matr1 := yummy.InputMatrix(rows, columns)
	fmt.Println("Upper triangle: ")
	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			if i <= j {
				fmt.Print(matr1[i][j], "\t")
				sum = sum + matr1[i][j]
			} else {
				fmt.Print("0", "\t")
			}
		}
		fmt.Println()
	}
	fmt.Println("Sum of the upper triangle: ", sum)
}
