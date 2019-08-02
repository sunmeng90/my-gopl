package main

import "fmt"

func main() {
	// using index to get non-ascii character from string will get error
	s := "hello 北京"
	fmt.Println("Loop over bytes of string")
	for i := 0; i < len(s); i++ { //loop over bytes
		fmt.Printf("%c \n", s[i]) // error for non-ascii characters
	}
	fmt.Println("Loop over runes(char) of string")
	rs := []rune(s)
	for i := 0; i < len(rs); i++ { //loop over bytes
		fmt.Printf("%c \n", rs[i]) // error for non-ascii characters
	}
	fmt.Println("Range over of string")
	for _, v := range s { // v is unicode code point
		fmt.Printf("%c \n", v) // %c show the character
	}
}
