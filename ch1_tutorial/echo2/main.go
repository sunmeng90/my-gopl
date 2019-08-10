package main

import (
	"fmt"
	"os"
)

// go run ch1_tutorial/echo2/main.go  a b c
func main() {
	var s, sep string
	for _, v := range os.Args[1:] {
		s += sep + v
		sep = " "
	}

	fmt.Println(s)
}
