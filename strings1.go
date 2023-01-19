package main

import (
	"fmt"
)

// func main() {
// 	reader := bufio.NewReader(os.Stdin)
// 	fmt.Println("Enter the string: ")
// 	strData, _ := reader.ReaderString('\n')
// 	for i := 0; i < len(strData); i++ {
// 		fmt.Printf("The ASCII value of %c = %d\n", strData[i], strData[i])
// 	}
// }

func main() {
	var str string
	fmt.Println("Enter the string: ")
	fmt.Scanln(&str)
	for i := 0; i < len(str); i++ {
		fmt.Printf("The ASCII value of %c = %d\n", str[i], str[i])
	}
}
