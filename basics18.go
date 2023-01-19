package main

import "fmt"

func main() {
	var num, res int
	fmt.Println("Enter the number: ")
	fmt.Scanln(&num)
	fmt.Println("The multiplicative table: ")
	for i := num - 1; i < num; i++ {
		for y := 1; y <= 10; y++ {
			res = num * y
			fmt.Println(num, "*", y, "=", res)
		}
	}
}
