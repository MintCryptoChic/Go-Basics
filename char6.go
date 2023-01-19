package main 

import ("fmt"
"bufio"
"os"
"unicode")

func main(){
	reader:=bufio.NewReader(os.Stdin)
	fmt.Println("Enter the letter you want to turn to uppercase: ")
	lw,_,_:= reader.ReadRune()
	if unicode.IsLetter(lw){
		up:=unicode.ToUpper(lw)
		fmt.Printf("The upper case of %c is %c\n", lw, up)
	} else{
		fmt.Printf("Please enter a valid alphabet\n")
	}
}