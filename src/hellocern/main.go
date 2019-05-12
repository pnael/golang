package main

import (
	"fmt"
	"github.com/golang/example/stringutil"
)

func main() {
	msg := stringutil.Reverse("Hello, CERN")
	fmt.Println(msg)
}
