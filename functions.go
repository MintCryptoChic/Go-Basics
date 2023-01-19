package yummy

import "fmt"

func InputArray(size int) (arr2 []int) {
	arr1 := make([]int, size)
	for i := 0; i < size; i++ {
		fmt.Scan(&arr1[i])
	}
	arr2 = arr1
	return
}

func InputMatrix(rows int, columns int) (matr2 [10][10]int) {
	var matr1 [10][10]int
	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			fmt.Scan(&matr1[i][j])
			matr2[i][j] = matr1[i][j]
		}
	}
	fmt.Println("Matrix: ")
	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			fmt.Print(matr2[i][j], "\t")
		}
		fmt.Println()
	}
	return
}
