package main

import (
	"fmt"
	"math"
)

func main() {
	var r, h, l, sa, vol, lsa float64
	fmt.Println("CONE: ")
	fmt.Println("Enter the radius: ")
	fmt.Scanln(&r)
	fmt.Println("Enter the height ")
	fmt.Scanln(&h)
	l = math.Sqrt(r*r + h*h)
	sa = math.Pi*r*l + math.Pi*r*r
	vol = (1.0 / 3) * math.Pi * r * r * h
	lsa = math.Pi * r * l
	fmt.Println("Length: ", l)
	fmt.Println("Surface area: ", sa)
	fmt.Println("Volume: ", vol)
	fmt.Println("Lateral surface area: ", lsa)
}
