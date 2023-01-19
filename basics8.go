package main

import "fmt"

func main() {
	var units, surcharge, totamt, amt float64
	fmt.Println("Enter the number of units consumed: ")
	fmt.Scanln(&units)
	if units < 50 {
		amt = units * 2.60
		surcharge = 25
	} else if units <= 100 {
		amt = 130 + ((units - 50) * 3.25)
		surcharge = 35
	} else if units <= 200 {
		amt = 130 + 162.50 + ((units - 100) * 5.26)
		surcharge = 45
	} else {
		amt = 130 + 162.50 + 526 + ((units - 200) * 7.75)
		surcharge = 55
	}
	totamt = surcharge + amt
	fmt.Println("Total bill: ", totamt)
}
