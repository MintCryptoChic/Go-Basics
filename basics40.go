package main

import "fmt"

func main() {
	var num, sum int
	fmt.Println("Enter the maximum number: ")
	fmt.Scanln(&num)
	for i := 1; i <= num; i++ {
		if i%2 == 0 {
			sum = sum + i
		}
	}
	fmt.Println("The of even numbers is: ", sum)
}
