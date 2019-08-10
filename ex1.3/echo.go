package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println(strings.Join(os.Args[1:], " "))
	end := time.Now()
	fmt.Printf("Time duration: %v\n", end.Sub(start).String())

	start = time.Now()
	var s, sep string
	for _, v := range os.Args[1:] {
		s += sep + v
		sep = " "
	}
	fmt.Println(s)
	end = time.Now()
	fmt.Printf("Time duration: %v\n", end.Sub(start).String())
}
