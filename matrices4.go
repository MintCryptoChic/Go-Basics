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
	count := 1
	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			if (i == j) && (matr1[i][j] != 1) {
				count = 0
			} else if (i != j) && (matr1[i][j] != 0) {
				count = 0
			}
		}
	}
	if count == 1 {
		fmt.Println("It is an identity matrix")
	} else {
		fmt.Println("It is not an identity matrix")
	}
}
