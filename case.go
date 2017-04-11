package main

import (
	"fmt"
	"runtime"
)



func main() {
	fmt.Print("Go run on ")
	switch os:=runtime.GOOS; os {
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Println("other")
	}

}

