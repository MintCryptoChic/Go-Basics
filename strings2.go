package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the first string to concat: ")
	str1, _ := reader.ReadString('\n')
	str1 = strings.TrimSuffix(str1, "\n")
	fmt.Println("Enter the second string to concat: ")
	str2, _ := reader.ReadString('\n')
	str2 = strings.TrimSuffix(str2, "\n")
	str3 := str1 + " " + str2
	fmt.Println(str3)
}
