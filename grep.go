package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Grep")

	flag1ptr := flag.String("flag1", "No", " First arg")
	flag.Parse()

	fmt.Println("Flag 1", *flag1ptr)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
}
