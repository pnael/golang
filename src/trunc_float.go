package main

import (
    "fmt"
    "strconv"
    )

func main() {
    var x string
    fmt.Println("Type a float: ")
    fmt.Scan(&x)
    value,err := strconv.ParseFloat(x,32)
    if err != nil {
        fmt.Println("Error ")
    }
    fmt.Printf("%.0f", value)
    }
