package main

func main() {
	println(comma("1"))
	println(comma("12"))
	println(comma("123"))
	println(comma("1234"))
	println(comma("12345"))
	println(comma("123456"))
}

// comma insert commas in a non-negative decimal integer string.
func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}
