package main

import (
	"basics/yummy"
	"fmt"
)

func main() {
	var rows, columns, num int
	fmt.Println("Enter rows: ")
	fmt.Scanln(&rows)
	fmt.Println("Enter columns: ")
	fmt.Scanln(&columns)
	fmt.Println("Enter the matrix elements: ")
	matr1 := yummy.InputMatrix(rows, columns)
	fmt.Println("Enter the number: ")
	fmt.Scanln(&num)
	fmt.Println("Matrix after multiplication: ")
	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			matr1[i][j] = num * matr1[i][j]
			fmt.Print(matr1[i][j], "\t")
		}
		fmt.Println()
	}
}
