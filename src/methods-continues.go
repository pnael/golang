package main

import (
	"fmt"
	"math"
)

type MyFloat float64

func (f MyFloat) Abs() MyFloat {
	if f < 0 {
		return (-f)
	}
	return f
}

func main() {
	f := MyFloat(-math.Sqrt(2))

	fmt.Println(f.Abs())
}
