package main

import (
	"fmt"
	"strings"
)

func main() {
	board := [][]string{
		[]string{"-", "-", "-"},
		[]string{"-", "-", "-"},
		[]string{"-", "-", "-"},
	}

	board[0][0] = "X"
	board[2][2] = "0"
	board[1][2] = "X"
	board[1][0] = "0"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}
