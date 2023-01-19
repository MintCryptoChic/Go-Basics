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
	up, _, _ := reader.ReadRune()
	if unicode.IsUpper(up) {
		fmt.Println("The entered character is uppercase")
	} else {
		fmt.Println("The entered character is not uppercase")
	}
}
