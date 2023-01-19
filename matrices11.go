package main

import (
	"basics/yummy"
	"fmt"
)

func main() {
	var rows, count int
	fmt.Println("Enter rows/ columns: ")
	fmt.Scanln(&rows)
	fmt.Println("Enter the matrix elements: ")
	matr1 := yummy.InputMatrix(rows, rows)
	for i := 0; i < rows; i++ {
		for j := 0; j < rows; j++ {
			if matr1[i][j] == 0 {
				count++
			}
		}
	}
	if count > (rows*rows)/2 {
		fmt.Println("It is a sparse matrix")
	} else {
		fmt.Println("It is not a sparse matrix")
	}
}
