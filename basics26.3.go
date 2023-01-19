package main

import "fmt"

func main() {
	var num int
	var min int
	var max int
	fmt.Println("Enter the minimum and maximum numbers respectively: ")
	fmt.Scanln(&min, &max)
	fmt.Println("The perfect numbers between ", min, "and", max, "are: ")
	for num = min; num <= max; num++ {
		var sum int
		for i := 1; i < num; i++ {
			if num%i == 0 {
				sum = sum + i
			}
		}
		if num == sum {
			fmt.Println(num)
		}
	}

}
