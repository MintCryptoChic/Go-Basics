package main

import "fmt"

func main() {
	var i, j, k, rows int
	fmt.Println("Enter the number of rows: ")
	fmt.Scanln(&rows)
	fmt.Println("Same Numbers on all Sides of a Square Pattern")
	for i = 1; i <= rows; i++ {
		for j = 1; j <= rows; j++ {
			if i < j {
				fmt.Printf("%d ", rows-i+1)
			} else {
				fmt.Printf("%d ", rows-j+1)
			}
		}
		for k = rows - 1; k > 0; k-- {
			if i < k {
				fmt.Printf("%d ", rows-i+1)
			} else {
				fmt.Printf("%d ", rows-k+1)
			}
		}
		fmt.Println()
	}
	for i = rows - 1; i > 1; i-- {
		for j = 1; j <= rows; j++ {
			if i < j {
				fmt.Printf("%d ", rows-i+1)
			} else {
				fmt.Printf("%d ", rows-j+1)
			}
		}
		for k = rows - 1; k > 0; k-- {
			if i < k {
				fmt.Printf("%d ", rows-i+1)
			} else {
				fmt.Printf("%d ", rows-k+1)
			}
		}
		fmt.Println()
	}
}
