package main

import (
	"fmt"
	"math"
)

func main() {
	var i, j, rows float64
	fmt.Println("Enter the number of rows: ")
	fmt.Scanln(&rows)
	fmt.Println("Exponentially Increased Star Pattern")
	for i = 0; i <= rows; i++ {
		for j = 1; j <= (math.Pow(2, i)); j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
}
