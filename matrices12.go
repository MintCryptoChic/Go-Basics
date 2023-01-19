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
	fmt.Println("Sum of the columns: ")
	for j := 0; j < columns; j++ {
		sum := 0
		for i := 0; i < rows; i++ {
			sum = sum + matr1[i][j]
		}
		fmt.Println("Column", j, ": ", sum)
	}
}
