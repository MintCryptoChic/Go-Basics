package main

import "fmt"

func main() {
	var num, sum int
	fmt.Println("Enter the number: ")
	fmt.Scanln(&num)
	for num >= 10 {
		for sum = 0; num > 0; num = num / 10 {
			sum = sum + (num % 10)
		}
		if sum >= 10 {
			num = sum
		} else {
			fmt.Println("The generic root of the number is: ", sum)
		}
	}
}
