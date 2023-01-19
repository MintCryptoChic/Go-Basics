package main

import (
	"bytes"
	"fmt"
)

func main() {
	byteArray := []byte{72, 101, 108, 108, 111, 32, 87, 111, 114, 108, 100}
	var str string
	// str = string(byteArray)
	str = bytes.NewBuffer(byteArray).String()
	fmt.Println(str)
}
