package main

import "fmt"

func count(x int) {
	notes := [8]int{500, 100, 50, 20, 10, 5, 2, 1}
	temp := x
	for i := 0; i < 8; i++ {
		fmt.Println(notes[i], "Notes: ", temp/notes[i])
		temp = temp % notes[i]
	}
}

func main() {
	var amt int
	fmt.Println("Enter the amount: ")
	fmt.Scanln(&amt)
	count(amt)

}
