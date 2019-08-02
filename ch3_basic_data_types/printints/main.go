package main

import (
	"bytes"
	"fmt"
)

func main() {
	println(intsToString([]int{}))
	println(intsToString([]int{1}))
	println(intsToString([]int{1, 2}))
	println(intsToString([]int{1, 2, 12, 3}))
}

func intsToString(values []int) string {
	var buf bytes.Buffer
	buf.WriteByte('[')
	for i, v := range values {
		if i > 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(&buf, "%d", v)
	}
	buf.WriteByte(']')
	return buf.String()
}
