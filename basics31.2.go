package main

import "fmt"

func main() {
	var num int
	fmt.Println("Enter the number: ")
	fmt.Scanln(&num)
	fmt.Println("The numbers from 1 to", num, "are: ")
	for i := 1; i <= num; i++ {
		fmt.Println(i)
	}
}
