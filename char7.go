package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the uppercase letter you want to change to lowercase: ")
	up, _, _ := reader.ReadRune()
	if unicode.IsLetter(up) {
		lw := unicode.ToLower(up)
		fmt.Printf("The lower case of %c is %c\n", up, lw)
	} else {
		fmt.Printf("Please enter a valid alphabet\n")
	}

}
