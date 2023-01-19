package main

import "fmt"

func calcFac(x int) int {
	if x == 0 || x == 1 {
		return 1
	} else {
		return x * calcFac(x-1)
	}
}

func main() {
	var num int
	fmt.Println("Enter the number: ")
	fmt.Scanln(&num)
	fmt.Println("The factorial is: ", calcFac(num))
}
