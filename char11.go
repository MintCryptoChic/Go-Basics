package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter a charcater to check for vowel or consonant: ")
	ch, _ := reader.ReadByte()
	if ch == 'a' || ch == 'A' || ch == 'e' || ch == 'E' || ch == 'o' || ch == 'O' || ch == 'u' || ch == 'U' {
		fmt.Printf("%c is a vowel\n", ch)
	} else {
		fmt.Printf("%c is a consonant\n", ch)
	}
}
