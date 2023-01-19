package main

import "fmt"

func prime(x int) int {
	var count int
	for i := 2; i < x/2; i++ {
		if x%i == 0 {
			count++
		}
	}
	return count
}

func main() {
	var num int
	fmt.Println("Enter the number: ")
	fmt.Scanln(&num)
	if prime(num) == 0 && num > 1 {
		fmt.Println("It is a prime number")
	} else {
		fmt.Println("It is not a prime number")
	}
}
