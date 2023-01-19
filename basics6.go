package main

import "fmt"

func main() {
	var num int
	fmt.Println("Enter the number: ")
	fmt.Scanln(&num)
	cube := num * num * num
	fmt.Println("The cube of the number: ", cube)
}
