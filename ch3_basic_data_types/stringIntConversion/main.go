package main

import (
	"fmt"
	"strconv"
)

// demo to show how to convert between string and integer
func main() {
	// int to string
	var str123A string = fmt.Sprintf("%d", 123)
	fmt.Println(str123A) //integer to string
	var str123B string = strconv.Itoa(123)
	fmt.Println(str123B) //integer to ascii

	// string to int
	int123, _ := strconv.Atoi("123")
	fmt.Println(int123)
}
