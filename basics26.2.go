package main

import "fmt"

func perfect(x int) int {
	var sum int
	for i := 1; i < x; i++ {
		if x%i == 0 {
			sum = sum + i
		}
	}
	return sum
}

func main() {
	var num int
	fmt.Println("Enter the number: ")
	fmt.Scanln(&num)
	if num == perfect(num) {
		fmt.Println("It is a perfect number")
	} else {
		fmt.Println("It is not a perfect number")
	}
}
