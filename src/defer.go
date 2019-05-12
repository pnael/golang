package main

import (
	"fmt"
)

func main() {
	defer fmt.Println("Hello")

	fmt.Println("world")
	for i := 0; i < 1000000000; i++ {
	}
	fmt.Println("world 2")
}
