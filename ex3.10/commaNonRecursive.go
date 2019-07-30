package main

import "bytes"

func main() {
	println(comma(""))
	println(comma("1"))
	println(comma("12"))
	println(comma("123"))
	println(comma("1234"))
	println(comma("12345"))
	println(comma("123456"))
	println(comma("1234567"))
	println(comma("12345678"))
	println(comma("123456789"))
	println(comma("1234567891"))
	println(comma("12345678910"))
}

func comma(s string) string {
	var buf bytes.Buffer
	n := len(s)
	re := n % 3
	for i := 1; i <= n; i++ {
		buf.WriteByte(s[i-1])
		if i != n && (i == re || (i-re)%3 == 0) {
			buf.WriteString(",")
		}
	}
	return buf.String()
}
