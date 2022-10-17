package main

import (
    "fmt"
    "strings"
    )

func main() {
    fmt.Println("Type a string")
    var answer string
    fmt.Scan(&answer)

    lower_ans := strings.ToLower(answer)
    fmt.Println(lower_ans)

    if strings.IndexAny(lower_ans, "i") == 0 {
        fmt.Println(lower_ans+" start with i")
        if strings.ContainsRune(lower_ans,'a') {
            fmt.Println(lower_ans+" contains an a")
            if strings.HasSuffix(lower_ans, "n") {
                fmt.Println(lower_ans+" finishes with an n")
            } else {
                fmt.Println("Not Found")
            }
        } else {
            fmt.Println("Not found")
        }

    } else {
        fmt.Println("Not found")
    }
        
    }