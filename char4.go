package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the character: ")
	lw, _, _ := reader.ReadRune()
	if unicode.IsLower(lw) {
		fmt.Println("The entered character is lowercase")
	} else {
		fmt.Println("The entered character is not lowercase")
	}
}
