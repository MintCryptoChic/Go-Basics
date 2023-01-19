package main

import "fmt"

func main() {
	var rows, i, j, num int
	fmt.Println("Enter the number of rows: ")
	fmt.Scanln(&rows)
	fmt.Println("Floyds Triangle")
	num = 1
	for i = 1; i <= rows; i++ {
		for j = 1; j <= i; j++ {
			fmt.Printf("%d ", num)
			num++
		}
		fmt.Println()
	}
}
