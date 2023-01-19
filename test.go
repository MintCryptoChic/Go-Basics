package main

import "fmt"

// func main() {
// 	var i, j, rows int
// 	fmt.Println("Enter the number of rows: ")
// 	fmt.Scanln(&rows)
// 	fmt.Println(" ")
// }

func hello(a int, b int) (z int, y int) {
	z = a + b
	y = a - b
	return
}

func main() {

	_, y := hello(6, 5)
	fmt.Println()
	fmt.Print(y, "\n")
}
