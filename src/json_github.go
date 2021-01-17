package main

import (
       fmt
       http
       )

func main() {
     resp, err  := http.Get("https://api.github.com/users/pnael")
     if err != nil {
       log.Fatal(err)
       }

     fmt.Println(resp)