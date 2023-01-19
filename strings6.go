package main

import "fmt"

func main() {
	var str string
	fmt.Println("Enter a string: ")
	fmt.Scanln(&str)
	lastchar := str[len(str)-1]
	fmt.Printf("First character in string: %c\n", lastchar)
}
