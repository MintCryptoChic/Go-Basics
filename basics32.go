package main

import "fmt"

func noloop(x int) {
	if x <= 100 {
		fmt.Println(x)
		noloop(x + 1)
	}
}

func main() {
	num := 1
	noloop(num)

}
