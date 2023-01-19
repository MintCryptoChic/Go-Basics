package main

import (
	"fmt"
	"math"
)

func main() {
	var num, exp float64
	fmt.Println("Enter the number: ")
	fmt.Scanln(&num)
	fmt.Println("Enter the power: ")
	fmt.Scanln(&exp)
	ans := math.Pow(num, exp)
	fmt.Println("The answer is: ", ans)
}
