package main

import (
	"fmt"
	"math"
)

func main() {
	var r, sa, vol, lsa float64
	fmt.Println("SPHERE: ")
	fmt.Println("Enter the radius: ")
	fmt.Scanln(&r)
	sa = 4 * math.Pi * r * r
	vol = (math.Sqrt(sa)) / (4 * math.Pi)
	lsa = 4 * math.Pi * r * r * r
	fmt.Println("Surface area: ", sa)
	fmt.Println("Volume: ", vol)
	fmt.Println("Lateral surface area: ", lsa)
}
