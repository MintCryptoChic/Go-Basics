package main

import "fmt"

func factorial(x int) int {
	var fact int
	fact = 1
	for i := 1; i <= x; i++ {
		fact = fact * i
	}
	return fact

}
func main() {
	var n, r int
	fmt.Println("Enter N: ")
	fmt.Scanln(&n)
	fmt.Println("Enter R: ")
	fmt.Scanln(&r)
	ncr := factorial(n) / (factorial(r) * factorial(n-r))
	fmt.Println("The nCr factorial is: ", ncr)

}
