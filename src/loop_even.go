package main

import "fmt"

func main() {
	for i, j := 0, 0; i < 1000; i, j = i+2, j+1 {
		fmt.Println(i, j)
	}
}
