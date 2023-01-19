package main

import "fmt"

func main() {
	var x, y int
	fmt.Println("Enter the first number 'x': ")
	fmt.Scanln(&x)
	fmt.Println("Enter the second number 'y': ")
	fmt.Scanln(&y)
	switch {
	case x > y:
		fmt.Println("Largest number 'x': ", x)
	case y > x:
		fmt.Println("Largest number 'y': ", y)
	default:
		fmt.Println("Both numbers are equal")
	}
}
