package main

import "fmt"

var x uint8 = 1

func main() {
	fmt.Println(x)
	x = 1 + 254
	fmt.Println(x)
	x = x + 1
	fmt.Println(x)
}
