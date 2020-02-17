package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"regexp"
)


func main() {
     file, err := os.Open("135-0.txt")
     if err != nil {
     	log.Fatal(err)
     }
     defer file.Close()

     scanner := bufio.NewScanner(file)

     for scanner.Scan() {
	line := scanner.Text()
	mots := strings.Split(line, " ")

	for _ , mot := range mots {
	    rege,err :=  regexp.Compile("[^a-zA-Z0-9]+")
	     if err != nil {
	        log.Fatal(err)
             }
	     regmot := rege.ReplaceAllString(mot, "")
	     fmt.Println(regmot)
	    }
     }
}
