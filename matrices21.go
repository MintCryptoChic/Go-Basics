package main

import (
	"basics/yummy"
	"fmt"
)

func main() {
	var rows, columns, count int
	fmt.Println("Enter rows: ")
	fmt.Scanln(&rows)
	fmt.Println("Enter columns: ")
	fmt.Scanln(&columns)
	fmt.Println("Enter the first matrix elements: ")
	matr1 := yummy.InputMatrix(rows, columns)
	fmt.Println("Enter the second matrix elements: ")
	matr2 := yummy.InputMatrix(rows, columns)
	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			if matr1[i][j] != matr2[i][j] {
				count++
			}
		}
	}
	if count > 0 {
		fmt.Println("The matrices are not equal")
	} else {
		fmt.Println("The matrices are equal")
	}
}
