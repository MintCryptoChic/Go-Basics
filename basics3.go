package main

import (
	"fmt"
	"math"
)

func main() {
	var Pamount, CompoundI, InterestRate, timePeriod, ciFuture float64
	fmt.Println("Enter the principal amount: ")
	fmt.Scanln(&Pamount)
	fmt.Println("Enter the interest rate: ")
	fmt.Scanln(&InterestRate)
	fmt.Println("Enter the numbe rof years: ")
	fmt.Scanln(&timePeriod)
	ciFuture = Pamount * (math.Pow((1 + InterestRate/100), timePeriod))
	CompoundI = ciFuture - Pamount
	fmt.Println("Future Compound Interest: ", ciFuture)
	fmt.Println("Compound Interest: ", CompoundI)
}
