package main

import "fmt"

func main() {
	var basicSal, hra, da, grossSal float64
	fmt.Println("Enter the basic salary of the employee: ")
	fmt.Scanln(&basicSal)
	fmt.Println("Enter the HRA percentage: ")
	fmt.Scanln(&hra)
	fmt.Println("Enter the DA percentage: ")
	fmt.Scanln(&da)
	h := basicSal * (hra / 100)
	d := basicSal * (da / 100)
	grossSal = h + d + basicSal
	fmt.Println("HRA: ", h)
	fmt.Println("DA: ", d)
	fmt.Println("Total salary: ", grossSal)
}
