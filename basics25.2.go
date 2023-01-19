package main

import "fmt"

func palindrome(x int) int {
	rev := 0
	for temp := x; temp > 0; temp = temp / 10 {
		rev = rev*10 + (temp % 10)
	}
	return rev
}

func main() {
	var num int
	fmt.Println("Enter the number: ")
	fmt.Scanln(&num)
	if palindrome(num) == num {
		fmt.Println("The number is a palindrome")
	} else {
		fmt.Println("The number is not a palindrome")
	}
}
