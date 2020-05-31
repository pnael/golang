package main

import (
  "fmt"
  "io/ioutil"
  "net/http"
  "os"
)

func main() {
  for _, url := range os.Args[1:] {
    resp, err := http.Get(url)
    if err  != nil {
      fmt.Printf(os.Stderr, "fetch: %v\n", err)
      os.Exit(1)
    }
  }
}
