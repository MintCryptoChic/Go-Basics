package main

import "fmt"

func main() {
	var num, sq int
	fmt.Println("Enter the number: ")
	fmt.Scanln(&num)
	sq = num * num
	fmt.Println("The square of the number is: ", sq)
}
