package main

import "fmt"

func main() {
	var min, max int
	fmt.Println("Enter the minimum number: ")
	fmt.Scanln(&min)
	fmt.Println("Enter the maximum number: ")
	fmt.Scanln(&max)
	fmt.Println("The odd numbers from ", min, "to ", max, "are")
	if min%2 == 0 {
		min++
	}
	for min <= max {
		fmt.Println(min)
		min = min + 2
	}
}
