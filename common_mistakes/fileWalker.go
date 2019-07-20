package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func main() {
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots =[]string{"."}
	}
	fmt.Printf("start walking: %v\n", roots)
	fc := make(chan os.FileInfo)
	// sum := *new(int64)
	sum := int64(0)
	go func(){
		for _, d := range roots{
			walk(d, fc)
		}
		close(fc)
	}()

	for file := range fc {
		// fmt.Printf("Got file %s, size: %d\n", file.Name(), file.Size())
		sum = sum + file.Size()
	}
	fmt.Printf("total file size: %d", sum)
}

func walk(dir string, fc chan<- os.FileInfo) {
	for _, fd := range dirents(dir) {
		// fmt.Printf("%s\n", fd.Name())
		if fd.IsDir() {
			walk(filepath.Join(dir, fd.Name()), fc)
		} else {
			fc <- fd
		}
	}
}

func dirents(dir string) []os.FileInfo {
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		// fmt.Fprintf(os.Stderr, "du1: %v\n", err)
		return nil
	}
	return entries
}
