package main

import "fmt"

const (
      PI float64  = 3.14
      A = iota
      B = iota
)

func main() {
     var message string
     var a, b, c int

     a =1
     message = "Hello World !"
     fmt.Println(message)
     fmt.Println(PI,A, B)
     fmt.Println(message, a, b ,c)
     }