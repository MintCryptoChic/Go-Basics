package main

import "fmt"

func main() {
	var min, max int
	fmt.Println("Enter the minimum number: ")
	fmt.Scanln(&min)
	fmt.Println("Enter the maximum number: ")
	fmt.Scanln(&max)

	if min%2 != 0 {
		min++
	}
	fmt.Println("The even numbers from", min, "to", max, "are: ")
	for i := min; i <= max; i = i + 2 {
		fmt.Println(i)
	}
}
