package main

import "fmt"

func main() {
	var num, sum int
	fmt.Println("Enter the number: ")
	fmt.Scanln(&num)
	for i := 1; i < num; i++ {
		if num%i == 0 {
			sum = sum + i
		}
	}
	if num == sum {
		fmt.Println("It is a perfect number")
	} else {
		fmt.Println("It is not a perfect number")
	}
}
