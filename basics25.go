package main

import "fmt"

func main() {
	var num, temp, rev int
	rev = 0
	fmt.Println("Enter the number: ")
	fmt.Scanln(&num)
	for temp = num; temp > 0; temp = temp / 10 {
		rev = rev*10 + (temp % 10)
	}
	if rev == num {
		fmt.Println("The number is a palindrome")
	} else {
		fmt.Println("The number is not a palindrome")
	}
}
