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
	fmt.Println("start walking")
	fc := make(chan os.FileInfo)
	
	go func(){
		walk(".", fc)
		close(fc)
	}()

	for file := range fc {
		fmt.Printf("Got file %s\n", file.Name())
	}

}

func walk(dir string, fc chan<- os.FileInfo) {
	for _, fd := range dirents(dir) {
		fmt.Printf("%s\n", fd.Name())
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
		fmt.Fprintf(os.Stderr, "du1: %v\n", err)
		return nil
	}
	return entries
	//fmt.Printf("%v", entries)
	//for _, fd := range entries {
	//	fmt.Printf("%s\n", fd.Name())
	//}
}
