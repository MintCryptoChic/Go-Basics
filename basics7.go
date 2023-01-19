package main

import "fmt"

func main() {
	var basicSal, hra, da, grossSal float64
	fmt.Println("Enter the basic salary of the employee: ")
	fmt.Scanln(&basicSal)
	if basicSal <= 10000 {
		hra = (basicSal * 8) / 100
		da = (basicSal * 10) / 100
	} else if basicSal <= 20000 {
		hra = (basicSal * 16) / 100
		da = (basicSal * 20) / 100
	} else {
		hra = (basicSal * 24) / 100
		da = (basicSal * 30) / 100
	}

	grossSal = basicSal + da + hra
	fmt.Println("Total salary of employee: ", grossSal)
}
