package main

import (
	"fmt"
	"math"
)

var (
	x int
	y int
)

func main() {
	x:=3
	y:=4

	var f float64 = math.Sqrt(float64(x*x+y*y))
	var j int = int(f)

	fmt.Println(x,y,f,j)
} 
