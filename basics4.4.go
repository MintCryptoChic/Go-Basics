package main

import "fmt"

var count int = 0

func digicount(num int) int {
	for num > 0 {
		count = count + 1
		digicount(num / 10)
	}
	return count
}

func main() {
	var num int
	fmt.Println("Enter the number: ")
	fmt.Scanln(&num)
	count = digicount(num)
	fmt.Println("The number of digits: ", count)
}
