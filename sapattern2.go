package main

import "fmt"

func main() {
	var i, j, k, rows int
	fmt.Println("Enter the number of rows: ")
	fmt.Scanln(&rows)
	fmt.Println("Hollow Diamond Star Pattern")
	for i = 1; i <= rows; i++ {
		for j = 1; j <= rows-i; j++ {
			fmt.Print(" ")
		}
		for k = 1; k <= i*2-1; k++ {
			if k == 1 || k == i*2-1 {
				fmt.Printf("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
	for i = rows - 1; i > 0; i-- {
		for j = 1; j <= rows-i; j++ {
			fmt.Print(" ")
		}
		for k = 1; k <= i*2-1; k++ {
			if k == 1 || k == i*2-1 {
				fmt.Printf("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}

}
