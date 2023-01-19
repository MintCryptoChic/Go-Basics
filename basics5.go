package main

import "fmt"

func main() {
	notes := [8]int{500, 100, 50, 20, 10, 5, 2, 1}
	var amt int
	fmt.Println("Enter the amount: ")
	fmt.Scanln(&amt)
	temp := amt
	for i := 0; i < 8; i++ {
		fmt.Println(notes[i], "Notes: ", temp/notes[i])
		temp = temp % notes[i]
	}
}
