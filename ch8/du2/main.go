package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

var verbose = flag.Bool("v", false, "show verbose progress message")

func main() {
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}
	fmt.Printf("start walking with progress: %v\n", roots)
	fc := make(chan os.FileInfo)
	totalSize := int64(0)
	nfiles := 0
	var tick <-chan time.Time
	if *verbose {
		tick = time.Tick(500 * time.Millisecond)
	}
	go func() {
		for _, d := range roots {
			walk(d, fc)
		}
		close(fc)
	}()
loop:
	for {
		select {
		case file, ok := <-fc:
			if !ok {
				break loop
			}
			nfiles++
			totalSize += file.Size()
		case <-tick:
			pringDiskUsage(nfiles, totalSize)
		}
	}
	pringDiskUsage(nfiles, totalSize)
}

func pringDiskUsage(nfiles int, totalSize int64) {
	fmt.Printf("%d files %.2f GB\n", nfiles, float64(totalSize)/1e9)
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
