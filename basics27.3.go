package main

import "fmt"

func main() {
	var min, max int
	fmt.Println("Enter the minimum and the maximum values between which you want to print all the prime numbers: ")
	fmt.Scanln(&min, &max)
	fmt.Println("The prime numbers between", min, "and", max, "are: ")
	for i := min; i <= max; i++ {
		var count int
		for j := 2; j < i/2; j++ {
			if i%j == 0 {
				count++
			}
		}
		if count == 0 {
			fmt.Println(i)
		}
	}

}
