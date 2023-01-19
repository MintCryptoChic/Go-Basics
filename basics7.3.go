package main

import "fmt"

func main() {
	var basicSal, hra, da, grossSal float64
	fmt.Println("Enter the basic salary: ")
	fmt.Scanln(&basicSal)
	fmt.Println("Enter the HRA: ")
	fmt.Scanln(&hra)
	fmt.Println("Enter the DA: ")
	fmt.Scanln(&da)
	grossSal = basicSal + hra + da
	fmt.Println("Total Salary: ", grossSal)
}
