package main

import "fmt"

func main() {
	var amt, totamt, surcharge, units float64
	fmt.Println("Enter the number of units consumed: ")
	fmt.Scanln(&units)
	if units > 500 {
		amt = units * 12.65
		surcharge = 125
	} else if units >= 300 {
		amt = units * 10.75
		surcharge = 100
	} else if units >= 200 {
		amt = units * 8.26
		surcharge = 85
	} else if units >= 100 {
		amt = units * 5.98
		surcharge = 65
	} else {
		amt = units * 3.85
		surcharge = 45
	}
	totamt = surcharge + amt
	fmt.Println("Total bill: ", totamt)
}
