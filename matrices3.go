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
	det := (matr1[0][0]*matr1[1][1] - matr1[0][1]*matr1[1][0])
	fmt.Println("Determinant: ", det)
}
