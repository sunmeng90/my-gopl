package main

import (
	"fmt"
	"os"
)

// Echo its arguments per line
// go run ex1.2/echoPerLine.go a b c
func main() {
	for _, v := range os.Args[1:] {
		fmt.Println(v)
	}
}
