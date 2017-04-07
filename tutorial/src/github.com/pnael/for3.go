package main

import "fmt"

func main() {
	for _, x := range "abcdefghijklmnopqrstuvwxyz" {
		fmt.Printf("%v\n", x)
	}
}
