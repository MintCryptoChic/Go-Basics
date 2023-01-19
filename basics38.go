package main

import "fmt"

func main() {
	var num, sum, avg, i float64
	fmt.Println("Enter the maximum number: ")
	fmt.Scanln(&num)
	for i = 1; i <= num; i++ {
		sum = sum + i
	}
	avg = sum / num
	fmt.Println("The sum is: ", sum)
	fmt.Println("The average is: ", avg)
}
