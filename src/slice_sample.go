package main

import (
	"fmt"
)

func main() {
	mySlice := []int{1, 2, 3, 4, 5}
	mySlice2 := mySlice[0:3]
	mySlice3 := mySlice[1:4]
	fmt.Println(mySlice, len(mySlice))
	fmt.Println(mySlice2)
	fmt.Println(mySlice3)

}
