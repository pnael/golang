package main

import "fmt"

func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "world"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

    other_array := [...]int{1,2,3,4}
    fmt.Println(other_array)

    
    for i, v := range other_array {
        fmt.Printf("Index %d Value %d",i,v)
        fmt.Println()
    }
}
