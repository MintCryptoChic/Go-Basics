package main

import "fmt"

func main() {
	var si, pa, y, roi float64
	fmt.Println("Enter the principle amount: ")
	fmt.Scanln(&pa)
	fmt.Println("Enter the rate of interest: ")
	fmt.Scanln(&roi)
	fmt.Println("Enter the number of years: ")
	fmt.Scanln(&y)
	si = (pa * roi * y) / 100
	fmt.Println("The simple interest is: ", si)
}
