package main

import (
	"fmt"
)

func isInSlice(x int, list []int) bool {
	for _, b := range list {
		if b == x {
			return true
		}
	}
	return false
}

func minfree(idlist []int) int {
	x := 0
	for i := 0; i < len(idlist); i++ {
		if isInSlice(x, idlist) {
			x++
		} else {
			return x
		}
	}
	return x
}

func main() {

	idlist := []int{18, 4, 8, 9, 16, 1, 14, 7, 19, 3, 0, 5, 2, 11, 6}
	fmt.Println(idlist)
	id := minfree(idlist)
	fmt.Println(id)
}
