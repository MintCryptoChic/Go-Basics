package main

import (
	"bufio"
	"fmt"
	"os"
)

// func main() {
// 	var str string
// 	fmt.Println("Enter a string: ")
// 	fmt.Scanln(&str)
// 	l := len(str)
// 	fmt.Println("Length of string:", l)
// }

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter a string: ")
	str, _ := reader.ReadString('\n')
	leng := 0
	for _, l := range str {
		fmt.Printf("%c", l)
		leng++
	}
	fmt.Println("Length: ", leng-1)

}
