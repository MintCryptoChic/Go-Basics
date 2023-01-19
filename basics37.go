package main

import (
	"fmt"
	"math"
)

func main() {
	var num, sqr float64
	fmt.Println("Enter the number: ")
	fmt.Scanln(&num)
	sqr = math.Sqrt(num)
	fmt.Println("The square root of the number is: ", sqr)
}
