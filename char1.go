package main

import ("fmt"
"bufio"
"os"
"unicode")

func main(){
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter any character to check alphabet: ")
	alr,_,_ := reader.ReadRune()
	if unicode.IsLetter(alr){
		fmt.Printf("%c is an alphatbet\n", alr)
	} else {
		fmt.Printf("%c is not an alphabet\n", alr)
	}
}