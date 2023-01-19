package main

import (
	"fmt"
)

func main() {
	var l, sa, vol, lsa float64
	fmt.Println("CUBE: ")
	fmt.Println("Enter the length: ")
	fmt.Scanln(&l)
	sa = 6 * l * l
	vol = l * l * l
	lsa = 4 * (l * l)
	fmt.Println("Surface area: ", sa)
	fmt.Println("Volume: ", vol)
	fmt.Println("Lateral surface area: ", lsa)
}
