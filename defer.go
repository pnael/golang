package main

import (
	"fmt"
)

func main() {
	defer fmt.Println("Hello")

	fmt.Println("world")
	fmt.Println("world 2")
}
