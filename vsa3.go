package main

import (
	"fmt"
)

func main() {
	var l, w, h, sa, vol, lsa float64
	fmt.Println("CUBOID: ")
	fmt.Println("Enter the length: ")
	fmt.Scanln(&l)
	fmt.Println("Enter the width: ")
	fmt.Scanln(&w)
	fmt.Println("Enter the height: ")
	fmt.Scanln(&h)
	sa = (2 * l * w) + (2 * l * h) + (2 * h * w)
	vol = l * w * h
	lsa = 2 * h * (l + w)
	fmt.Println("Surface area: ", sa)
	fmt.Println("Volume: ", vol)
	fmt.Println("Lateral surface area: ", lsa)
}
