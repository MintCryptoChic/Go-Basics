package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter any character to find ASCII: ")
	ch, _ := reader.ReadByte()
	fmt.Printf("The ASCII value of %c = %d\n", ch, ch)
}
