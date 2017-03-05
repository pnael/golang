package main

import "fmt"

func half(x int) (int, bool) {
	even := false
	if x%2 == 0 {
		even = true
	}
	return x / 2, even
}

func add(x ...int) int {
	max := x[0]
	for _, v := range x {
		if v > max {
			max = v
		}
	}
	return max
}

func main() {
	fmt.Println(half(2))
	fmt.Println(half(1))

	fmt.Println(add(1, 2, 8, 9, 7, 6, 4))
}
