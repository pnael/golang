package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

func ispermutation(tab string, tab2 string, debug bool) bool {
	pareil := false
	if debug {
		fmt.Println("Debug :", tab, tab2)
	}

	for _, d := range tab2 {
		for _, c := range tab {
			fmt.Println("d,c:", d, c)
			if c == d {
				pareil = true
			} else {
				pareil = false
			}
		}
	}

	if len(tab) != len(tab2) {
		return false
	}
	return pareil
	return true
}

var debug bool

func main() {

	if len(os.Args) == 3 {
		debug = true
	}
	thisdigit := os.Args[1]

	newdigit, err := strconv.Atoi(thisdigit)
	if err != nil {
		log.Fatal(err)
	}

	newdigit++

	if debug {
		fmt.Println(debug, newdigit, thisdigit)
		time.Sleep(5)
	}

	for {
		newdigit++
		newdigitstring := strconv.Itoa(newdigit)
		isFinished := ispermutation(newdigitstring, thisdigit, debug)

		if newdigitstring == "1254" {
			fmt.Println("Fini", isFinished)
			os.Exit(0)
		}
	}

	fmt.Println(newdigit)
}
