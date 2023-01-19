package main

import "fmt"

func main() {
	var i, j, rows, num int
	fmt.Println("Enter the number of rows: ")
	fmt.Scanln(&rows)
	fmt.Println("Right Triangle of Consecutive Rows Pattern")
	for i = 1; i <= rows; i++ {
		num = i
		for j = 1; j <= i; j++ {
			fmt.Printf("%d ", num)
			num = num + rows - j
		}
		fmt.Println()
	}
}
