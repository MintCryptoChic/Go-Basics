package main

import "fmt"

func main() {
	var min, max int
	fmt.Println("Enter the minimum number: ")
	fmt.Scanln(&min)
	fmt.Println("Enter the maximum value: ")
	fmt.Scanln(&max)
	fmt.Println("The odd numbers are: ")
	for i := min; i <= max; i++ {
		if i%2 != 0 {
			fmt.Println(i)
		}
	}
}
