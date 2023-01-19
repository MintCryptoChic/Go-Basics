package main

import "fmt"

func main() {
	var num, exp, ans, i float64
	ans = 1
	fmt.Println("Enter the number: ")
	fmt.Scanln(&num)
	fmt.Println("Enter the power: ")
	fmt.Scanln(&exp)
	for i = 1; i <= exp; i++ {
		ans = ans * num
	}
	fmt.Println("The answer is: ", ans)
}
