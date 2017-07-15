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
	start := 0

	for scanner.Scan() {
		if start == 0 {
			start = 1
			continue
		}
		current := scanner.Text()
		nb1 := strings.Count(current, "-+")
		nb2 := strings.Count(current, "+-")
		height := 1 + nb1 + nb2
		lastc := current[len(current)-1]

		//fmt.Println(current, ", -+ :", nb1, ", +-:", nb2, ", height:", height, ", last char:", lastc)

		if lastc == 45 {
			fmt.Printf("Case #%d: %d\n", start, height)
		} else {
			fmt.Printf("Case #%d: %d\n", start, height-1)
		}
		start++
	}
}
