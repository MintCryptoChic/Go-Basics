package main

import "fmt"

func main() {
	var num, sum int
	fmt.Println("Enter the number: ")
	fmt.Scanln(&num)
	for sum = 0; num > 0; num = num / 10 {
		sum = sum + num%10
	}
	fmt.Println("The sum is: ", sum)
}
