package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)



func main() {

	if len(os.Args) != 2 {
		log.Fatalln("Need one file name")
	}
	f, err := os.Open(os.Args[1])

	if err != nil {
		log.Fatalln("Can't open file", os.Args[1])
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	//Skip the number of case
  scanner.Scan()
	i:=0
	for scanner.Scan() {
    lineofword := scanner.Text()

		//Get all the words
		words := strings.Split(lineofword, " ")
    i++

    var answer string

    for j := len(words)-1; j>=0; j-- {
      answer = answer + " " + words[j]
    }

    fmt.Printf("Case #%d: %s\n",i,answer)

		}
}
