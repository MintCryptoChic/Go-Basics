package main

import "fmt"

func cube(x int) int {
	c := x * x * x
	return c
}

func main() {
	var num int
	fmt.Println("Enter the number: ")
	fmt.Scanln(&num)
	fmt.Println("The cube of the number: ", cube(num))
}
