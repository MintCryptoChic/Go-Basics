package main

import "fmt"

func factors(x int) {
	for i := 1; i <= x; i++ {
		if x%i == 0 {
			fmt.Println(i)
		}
	}

}

func main() {
	var x int
	fmt.Println("Enter the number: ")
	fmt.Scanln(&x)
	fmt.Println("The factor of the number are: ")
	factors(x)

}
