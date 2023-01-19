package main

import "fmt"

func count(x int) int {
	var c int
	for c = 0; x > 0; x = x / 10 {
		c = c + 1
	}
	return c
}
func main() {
	var num int
	fmt.Println("Enter the number: ")
	fmt.Scanln(&num)
	fmt.Println("Number of digits in the number: ", count(num))

}
