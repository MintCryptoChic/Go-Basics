package main

import "fmt"

func main() {
	var num int
	fmt.Println("Enter the number: ")
	fmt.Scanln(&num)
	if num%5 == 0 && num%11 == 0 {
		fmt.Println(num, "is divisible by 5 and 11")
	} else {
		fmt.Println(num, "is not divisible by 5 and 11")
	}
}
