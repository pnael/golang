package main

import (
	"flag"
	"fmt"
)

func main() {
	fmt.Println("Grep")

	flag1ptr := flag.String("flag1", "No", " First arg")
	flag.Parse()

	fmt.Println("Flag 1", *flag1ptr)
}
