package main

import "fmt"

func main() {
	var rows, i, j, k int
	fmt.Println("Enter the number of rows: ")
	fmt.Scanln(&rows)
	fmt.Println("Diamond Pattern")
	for i = 1; i <= rows; i++ {
		for j = 1; j <= rows-i; j++ {
			fmt.Print(" ")
		}
		for k = 1; k <= i*2-1; k++ {
			fmt.Printf("%d", k)
		}
		fmt.Println()
	}
	for i = rows - 1; i >= 1; i-- {
		for j = 1; j <= rows-i; j++ {
			fmt.Print(" ")
		}
		for k = 1; k <= i*2-1; k++ {
			fmt.Printf("%d", k)
		}
		fmt.Println()
	}

}
