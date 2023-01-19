package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the character to check for a digit: ")
	num, _, _ := reader.ReadRune()
	if unicode.IsDigit(num) {
		fmt.Printf("%c is a digit\n", num)
	} else {
		fmt.Printf("%c is not a digit\n", num)
	}
}
