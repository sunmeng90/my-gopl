package main

import "fmt"
import "strings"

func main() {
	pathA := basename("a")
	if pathA != "a" {
		fmt.Println(fmt.Errorf("basename for %s should be %s", "a", "a"))
	}
	pathB := basename("a.go")
	if pathB != "a" {
		fmt.Println(fmt.Errorf("basename for %s should be %s", "a.go", "a"))
	}
	pathC := basename("a/b/c.go")
	if pathC != "c" {
		fmt.Println(fmt.Errorf("basename for %s should be %s", "a/b/c.go", "c"))
	}
	pathD := basename("a/b.c.go")
	if pathD != "b.c" {
		fmt.Println(fmt.Errorf("basename for %s should be %s", "a/b.c.go", "b.c"))
	}
}

// basename removes directory components and a .suffix
// e.g., a => a, a.go => a, a/b/c.go => c, a/b.c.go => b.c
func basename(s string) string {
	lastIndex := strings.LastIndex(s, "/")
	s = s[lastIndex+1:]

	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}
	return s
}
