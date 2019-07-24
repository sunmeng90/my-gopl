//du4 calculate total size concurrntly with opening limited file handles
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var verbose = flag.Bool("v", false, "show verbose progress message")

var done = make(chan struct{})

func cancelled() bool {
	select {
	case <-done:
		return true
	default:
		return false
	}
}

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
		os.Stdin.Read(make([]byte, 1)) //wait for cancel signal from console
		close(done)
	}()
	var wg sync.WaitGroup

	for _, d := range roots {
		wg.Add(1)
		go walk(d, &wg, fc)
	}
	go func() {
		wg.Wait()
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
		case <-done:
			println("abort...")
			// before leaving, drain all the file from channel that other goroutines had already put into
			for range fc {
				//do nothing
			}
			return
		}
	}
	pringDiskUsage(nfiles, totalSize)
}

func pringDiskUsage(nfiles int, totalSize int64) {
	fmt.Printf("%d files %.2f GB\n", nfiles, float64(totalSize)/1e9)
}

func walk(dir string, n *sync.WaitGroup, fc chan<- os.FileInfo) {
	defer n.Done()
	if cancelled() {
		//pre-check for cancel flag and return if user request to cancel
		return
	}
	for _, fd := range dirents(dir) {
		if cancelled() {
			return
		}
		if fd.IsDir() {
			n.Add(1)
			go walk(filepath.Join(dir, fd.Name()), n, fc)
		} else {
			fc <- fd
		}
	}
}

var sema = make(chan struct{}, 2) //semaphore of size 20

func dirents(dir string) []os.FileInfo {
	select {
	case sema <- struct{}{}: //acquire a token
		defer func() { <-sema }()
	case <-done:
		return nil
	}
	// defer func() { <-sema }() //release token
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil
	}
	return entries
}
