package main

import "fmt"

func main() {
	var x, y int
	fmt.Println("Enter the first number 'x': ")
	fmt.Scanln(&x)
	fmt.Println("Enter the second number 'y': ")
	fmt.Scanln(&y)
	if x > y {
		fmt.Println("Largest number 'x': ", x)
	} else if x < y {
		fmt.Println("Largest number 'y': ", y)
	}
}
