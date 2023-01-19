package main

import "fmt"

func main() {
	var i, j, rows, num int
	fmt.Println("Enter the number of rows: ")
	fmt.Scanln(&rows)
	fmt.Println("Right Triangle of Consecutive Column Numbers")
	num = 1
	for i = 1; i <= rows; i++ {
		for j = 1; j <= i; j++ {
			fmt.Print(num, " ")
			num++
		}
		fmt.Println()
	}
}
