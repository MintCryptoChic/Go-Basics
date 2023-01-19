package main

import (
	"fmt"
	"math"
)

func main() {
	var r, h, sa, vol, lsa, tbsa float64
	fmt.Println("CYLINDER: ")
	fmt.Println("Enter the radius: ")
	fmt.Scanln(&r)
	fmt.Println("Enter the height: ")
	fmt.Scanln(&h)
	sa = 2*math.Pi*r*r + 2*math.Pi*r*h
	vol = math.Pi * r * r * h
	lsa = 2 * math.Pi * r * h
	tbsa = math.Pi * r * r
	fmt.Println("Surface area: ", sa)
	fmt.Println("Volume: ", vol)
	fmt.Println("Lateral surface area: ", lsa)
	fmt.Println("Top/bottom surface area: ", tbsa)
}
