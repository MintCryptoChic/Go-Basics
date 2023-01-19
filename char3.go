package main

import "fmt"

func main() {
	for i := 0; i <= 255; i++ {
		fmt.Printf("The ASCII value of %c = %d\n", i, i)
	}
}
