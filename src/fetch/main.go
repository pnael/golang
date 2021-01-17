// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 16.
//!+

// Fetch prints the content found at each specified URL.
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"flag"
)

func main() {
	printBody := flag.Bool("body",false, "Print the html page body")
	flag.Parse()
	//leftArgs := flag.Args()
	for i, url := range os.Args[1:] {
		fmt.Println(i)
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("Headers: %s",resp.Header)
		b, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		if *printBody {
			fmt.Printf("%s", b)
		}
	}
}

//!-
