package main

import "fmt"

func main() {
    sli := make([]int, 10)
    sli = append(sli,100)
    
    for _,v := range sli {
        fmt.Println(v)
    }
    }