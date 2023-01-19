package main

import (
	"basics/yummy"
	"fmt"
)

func main() {
	var rows, columns, temp int
	fmt.Println("Enter rows: ")
	fmt.Scanln(&rows)
	fmt.Println("Enter columns: ")
	fmt.Scanln(&columns)
	fmt.Println("Enter the matrix elements: ")
	matr1 := yummy.InputMatrix(rows, columns)
	if rows == columns {
		for i := 0; i < rows; i++ {
			temp = matr1[i][i]
			matr1[i][i] = matr1[i][rows-i-1]
			matr1[i][rows-i-1] = temp
		}
	} else {
		fmt.Println("The matrix is not a square matrix")
	}
	fmt.Println("The matrix with reversed diagonals: ")
	if rows == columns {
		for i := 0; i < rows; i++ {
			for j := 0; j < columns; j++ {
				fmt.Print(matr1[i][j], "\t")
			}
			fmt.Println()
		}
	}
}
