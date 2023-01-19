package main

import "fmt"

func main() {
	var num, count int
	fmt.Println("Enter the number: ")
	fmt.Scanln(&num)
	for i := 2; i < num/2; i++ {
		if num%i == 0 {
			count++
		}
	}
	if count == 0 && num > 1 {
		fmt.Println("It is a prime number")
	} else {
		fmt.Println("It is not a prime number")
	}
}
