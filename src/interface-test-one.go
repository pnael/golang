package main

import (
	"fmt"
)

type Abser interface {
	Abs() float64
}

func main() {
	var w Writer = ConsoleWriter{}
	w.Write("Hello Go !")
}

type Writer interface {
	Write([]byte) (int, error)
}

type ConsoleWriter struct{}

func (cw ConsoleWriter) Write(data []byte) (int, error) {
	n, err := fmt.Println(data)
	return n, err
}
