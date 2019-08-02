package main

import "fmt"

func main() {
	var greeting string = "hello world"
	r := []rune(greeting)
	for _, c := range r {
		fmt.Printf("%s\t->\t%d\n", string(c), c)
	}
	// convert a integer value to a string interprets the integer as a rune value, and yields the UTF-8 representation of that rune
	fmt.Printf("string(97)==\"a\"\t%t\n", string(97) == "a")
	fmt.Printf("string(97)==\"97\"\t%t\n", string(97) == "97")
}
