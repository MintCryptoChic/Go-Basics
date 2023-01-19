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
	fmt.Println("Enter the first matrix: ")
	matr1 := yummy.InputMatrix(rows, columns)
	fmt.Println("Enter the second matrix: ")
	matr2 := yummy.InputMatrix(rows, columns)
	fmt.Println("Matrix after multiplication: ")
	var matmul [10][10]int
	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			matmul[i][j] = matr1[i][j] * matr2[i][j]
			fmt.Print(matmul[i][j], "\t")
		}
		fmt.Println()
	}
}
