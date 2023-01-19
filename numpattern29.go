package main

import "fmt"

func main() {
	var i, j, k, rows int
	fmt.Println("Enter the number of rows: ")
	fmt.Scanln(&rows)
	fmt.Println("Sandglass Number Pattern")
	for i = 1; i <= rows; i++ {
		for j = 1; j < i; j++ {
			fmt.Print(" ")
		}
		for k = i; k <= rows; k++ {
			fmt.Printf("%d ", k)
		}
		fmt.Println()
	}
	for i = rows - 1; i > 0; i-- {
		for j = 1; j < i; j++ {
			fmt.Print(" ")
		}
		for k = i; k <= rows; k++ {
			fmt.Printf("%d ", k)
		}
		fmt.Println()
	}

}
