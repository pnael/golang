package main

import (
	"fmt"
	"math/rand"
)

func main() {
	if x := rand.Int(); x%2 == 0 {
		fmt.Println(x)
		fmt.Println("X est pair")
	} else {
		fmt.Println(x)
		fmt.Println("X est impair")
	}
}
