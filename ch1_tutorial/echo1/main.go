package main

import (
	"fmt"
	"os"
)

// Echo1 prints its command-line arguments separated by a blank character
// Usage: go run ch1_tutorial/echo1/main.go  a b c
func main() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}

	fmt.Println(s)
}
