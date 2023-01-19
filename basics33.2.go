package main

import "fmt"

func rev(x int) int {
	var reve int
	for x > 0 {
		reve = reve*10 + x%10
		x = x / 10
	}
	return reve
}
func main() {
	var num int
	fmt.Println("Enter the number: ")
	fmt.Scanln(&num)
	fmt.Println("The reverse of the number is: ", rev(num))
}
