package main

import (
	"basics/yummy"
	"fmt"
)

func main() {
	var rows, sum int
	fmt.Println("Enter rows/columns: ")
	fmt.Scanln(&rows)
	fmt.Println("Enter the matrix elements: ")
	matr1 := yummy.InputMatrix(rows, rows)
	for i := 0; i < rows; i++ {
		sum = sum + matr1[i][i]
	}
	fmt.Println("Sum of the diagonals: ", sum)
}
