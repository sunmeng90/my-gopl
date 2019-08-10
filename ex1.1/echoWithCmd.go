package main

import (
	"fmt"
	"os"
	"strings"
)

// Echo print command and its arguments
// go run ex1.1/echoWithCmd.go a b c
func main() {
	fmt.Println(strings.Join(os.Args, " "))
}
