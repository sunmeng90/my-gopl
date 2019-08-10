package main

import (
	"os"
	"strings"
)

// Usage: go run ch1_tutorial/echo3/main.go  a b c
func main() {
	//fmt.Println(strings.Join(os.Args[1:], " "))
	println(strings.Join(os.Args[1:], " "))
}
