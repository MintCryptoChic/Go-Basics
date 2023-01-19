package main

import "fmt"

func main() {
	var x int
	fmt.Println("Enter the number: ")
	fmt.Scanln(&x)
	fmt.Println("The factors of the number are: ")
	for i := 1; i <= x; i++ {
		if x%i == 0 {
			fmt.Println(i)
		}
	}
}
