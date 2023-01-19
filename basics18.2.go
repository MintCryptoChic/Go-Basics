package main

import "fmt"

func main() {
	var x, y int
	fmt.Println("Enter the first number: ")
	fmt.Scanln(&x)
	fmt.Println("Enter the second number: ")
	fmt.Scanln(&y)
	for i := x; i <= y; i++ {
		for j := 1; j <= 10; j++ {
			fmt.Println(i, "*", j, "=", i*j)
		}
		fmt.Println("------------------")
	}
}
