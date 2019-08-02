package main

import "fmt"

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
	for i := len(s) - 1; i > 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}

	for i := len(s) - 1; i > 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}
	return s
}
