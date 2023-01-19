package main

import "fmt"

func main() {
	var s, p, amt int
	fmt.Println("Enter the actual cost: ")
	fmt.Scanln(&p)
	fmt.Println("Enter the sale price: ")
	fmt.Scanln(&s)
	if s-p > 0 {
		amt = s - p
		fmt.Println("Total profit: ", amt)
	} else if s-p < 0 {
		amt = p - s
		fmt.Println("Total loss: ", amt)
	} else {
		fmt.Println("No profit no loss")
	}
}
