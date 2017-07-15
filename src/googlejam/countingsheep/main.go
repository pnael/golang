package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
  "strconv"
)


func CountSheep(cas, N int) {
  digit := make(map[int]int)
  var i int
  for i=1;;i++ {
    if N == 0 {
      fmt.Printf("Case #%d: INSOMNIA\n",cas)
      return
    }
    answer := strconv.Itoa(N*i)
    for _,c := range(answer) {
      digit[int(c)-48]++
    }

    if len(digit) == 10 {
      break
    }

  }
  fmt.Printf("Case #%d: %d\n", cas, N*i)

  return
}

func main() {

	if len(os.Args) != 2 {
		log.Fatalln("Need one file name")
	}
	f, err := os.Open(os.Args[1])

	if err != nil {
		log.Fatalln("Can't open file", os.Args[1])
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
  start := 0


	for scanner.Scan() {
    if start == 0 {
      start = 1
      continue
    }
		N,_ := strconv.Atoi(scanner.Text())
    CountSheep(start,N)
    start++

	}
}
