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
	totalSize := int64(0)
	totalFile := 0
	go func(){
		for _, d := range roots{
			walk(d, fc)
		}
		close(fc)
	}()

	for file := range fc {
		// fmt.Printf("Got file %s, size: %d\n", file.Name(), file.Size())
		totalSize += file.Size()
		totalFile++
	}
	fmt.Printf("%d files %.2f GB\n", totalFile, float64(totalSize)/1e9)
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
