package main

import "fmt"

func main() {
	var x int
	fmt.Println("Enter the natural number: ")
	fmt.Scanln(&x)
	fmt.Println("Natural Numbers are: ")
	for i := 1; i <= x; i++ {
		fmt.Println(i)
	}
}
