package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)


func main() {

  T9alpha := map[rune]string {
    'a':"2", 'b':"22", 'c':"222",
    'd':"3", 'e':"33", 'f':"333",
    'g':"4", 'h':"44", 'i':"444",
    'j':"5", 'k':"55", 'l':"555",
    'm':"6", 'n':"66", 'o':"666",
    'p':"7", 'q':"77", 'r':"777", 's':"7777",
    't':"8", 'u':"88", 'v':"888",
    'w':"9", 'x':"99", 'y':"999", 'z':"9999",
    ' ':"0",
  }


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
	i :=0
	for scanner.Scan() {
    lineofword := scanner.Text()
    fmt.Printf("Case #%d: ",i+1)
		//Get all the words
    var oldc rune
    for j, c := range lineofword {
      if j==0 {
        fmt.Printf("%s", T9alpha[c])
        oldc = rune(T9alpha[c][0])
        continue
      }
      if oldc ==  rune(T9alpha[c][0]) {
        fmt.Printf(" ")
      }

      oldc = rune(T9alpha[c][0])
      fmt.Printf("%s", T9alpha[c])
    }
    fmt.Println()
    i++


		}
}
