package main

import "fmt"

func main() {
	var num int
	fmt.Println("Enter the maximum number that may be divisible by 5 and 11")
	fmt.Scanln(&num)
	fmt.Println("Numbers divisible by 5 and 11 are: ")
	for num > 0 {
		if num%5 == 0 && num%11 == 0 {
			fmt.Println(num)
		}
		num--
	}
}
