package main

import "fmt"

func main() {
	var str string
	fmt.Println("Enter a string: ")
	fmt.Scanln(&str)
	fmt.Printf("%s", str)
	// fmt.Println(str)
}
