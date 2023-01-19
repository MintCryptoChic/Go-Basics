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
	var matr2 [10][10]int
	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			matr2[j][i] = matr1[i][j]
		}
	}
	fmt.Println("Transpose: ")
	for i := 0; i < columns; i++ {
		for j := 0; j < rows; j++ {
			fmt.Print(matr2[i][j], "\t")
		}
		fmt.Println()
	}
}
