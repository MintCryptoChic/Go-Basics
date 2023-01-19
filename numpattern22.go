package main

import "fmt"

func fact(num int) int {
	if num == 0 || num == 1 {
		return 1
	}
	return num * fact(num-1)
}

func main() {
	var i, j, num int
	fmt.Println("Enter the number of rows: ")
	fmt.Scanln(&num)
	fmt.Println("Pascals Triangle Pattern")
	for i = 0; i < num; i++ {
		for j = 1; j <= num-i-2; j++ {
			fmt.Printf(" ")
		}
		for j = 0; j <= i; j++ {
			fmt.Printf("%d ", fact(i)/(fact(j)*fact(i-j)))
		}
		fmt.Println()
	}
}
