package main

import "fmt"

func main() {
	var num, evensum, oddsum int
	fmt.Println("Enter the maximum number: ")
	fmt.Scanln(&num)
	for i := 1; i <= num; i++ {
		if i%2 == 0 {
			evensum = evensum + i
		} else if i%2 != 0 {
			oddsum = oddsum + i
		}
	}
	fmt.Println("The of even numbers is: ", evensum)
	fmt.Println("The of odd numbers is: ", oddsum)
}
