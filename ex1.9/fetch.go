package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Fetch: %v", err)
			os.Exit(-1)
		}
		_, err = io.Copy(os.Stdout, resp.Body)
		fmt.Fprintln(os.Stderr, resp.StatusCode)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Fetch: reading %s: %v\n", url, err)
			os.Exit(-1)
		}
	}
}
