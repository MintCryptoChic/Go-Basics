package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the starting character: ")
	st, _ := reader.ReadByte()
	fmt.Println("The alphabets are: ")
	for ch := st; ch <= 'Z'; ch++ {
		fmt.Printf("%c ", ch)
	}
	fmt.Println()
}

// func main() {
// 	fmt.Println("The alphabets are: ")
// 	for ch := 'A'; ch <= 'Z'; ch++ {
// 		fmt.Printf("%c \n", ch)
// 	}
// }
