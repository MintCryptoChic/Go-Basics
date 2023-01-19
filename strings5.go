package main

import "fmt"

func main() {
	var str string
	fmt.Println("Enter a string: ")
	fmt.Scanln(&str)
	firstchar := str[0]
	fmt.Printf("First character in string: %c\n", firstchar)
}
