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
	case "Windows"::
		fmt.Println("Windows")
	default:
		fmt.Println("other")
	}

}
