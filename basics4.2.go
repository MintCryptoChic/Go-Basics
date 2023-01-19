package main

import "fmt"

func main() {
	var num, count int
	count = 0
	fmt.Println("Enter the number: ")
	fmt.Scanln(&num)
	for count = 0; num > 0; num = num / 10 {
		count = count + 1
	}
	fmt.Println("The number of digits in the number: ", count)

}
