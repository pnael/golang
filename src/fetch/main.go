// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 16.
//!+

// Fetch prints the content found at each specified URL.
package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
)

func main() {
	printBody := flag.Bool("body", false, "Print the html page body")
	url := flag.String("url", "", "URL to reach")
<<<<<<< HEAD
	flag.Parse()
	//leftArgs := flag.Args()
=======
	printHeaders := flag.Bool("headers", false, "Print HTTP headers")
	flag.Parse()
>>>>>>> a06e1fae6ef76a9bf226f3419772991a1119de7b
	fmt.Println(*url)
	resp, err := http.Get(*url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		os.Exit(1)
	}

<<<<<<< HEAD
	fmt.Println("Headers: %s", resp.Header)
	b, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
		os.Exit(1)
	}

	if *printBody {
		fmt.Printf("%s", b)
=======
	if *printHeaders {
		fmt.Println("Headers:")
		for header, value := range resp.Header {
			fmt.Println(header, value)
		}
>>>>>>> a06e1fae6ef76a9bf226f3419772991a1119de7b
	}

	if *printBody {
		fmt.Println("Body:")
	}

}

//!-
