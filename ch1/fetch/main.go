// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 16.
//!+

// Fetch prints the content found at each specified URL.
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)
//book task 1.7-1.9
func main() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http://"){
			url = "http://"+url
		}

		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		fmt.Fprintf(os.Stderr, "status code: %v\n", resp.Status)

		_, err = io.Copy(os.Stderr, resp.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "copy: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		resp.Body.Close()


		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
	}
}

//!-
