package main

import "fmt"

func main() {
	var str string
	fmt.Println("Enter string: ")
	fmt.Scanln(&str)
	byteString := []byte(str)
	fmt.Println(byteString)

}
