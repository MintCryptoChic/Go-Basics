package main

import "fmt"

func genroot(x int) int {
	var sum int
	for sum = 0; x > 0; x = x / 10 {
		sum = sum + (x % 10)
	}
	return sum
}

func main() {
	var num, sum int
	fmt.Println("Enter the number: ")
	fmt.Scanln(&num)
	for num >= 10 {
		sum = genroot(num)
		if sum >= 10 {
			num = sum
		} else {
			fmt.Println("The generic root of the number is: ", sum)
			break
		}
	}
}
