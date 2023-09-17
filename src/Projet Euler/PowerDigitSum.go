package main

import (
	"fmt"
	"math"
)

func main() {
	var BigNumber int64
	BigNumber = int64(math.Pow(2, 1000))

	fmt.Println(BigNumber)
}
