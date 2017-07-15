package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func powerOf10(N int) uint64 {
	return uint64(math.Pow(10, float64(N)))
}

func printCoinJams(N, J int) {
	// Print to stdout the coinjam candidate
	lower := powerOf10(N)
	limit := uint64(1)

	for i := 1; i <= N; i++ {
		limit = limit + powerOf10(i)
	}

	cas := lower
	fmt.Printf("%d, %d\n", cas, limit)
	for i := 0; cas < uint64(limit); i++ {
		value := uint64(0)
		for j := i; j > 0; j-- {
			value = value + powerOf10(j)
			cas = lower + value + 1
			for base := 2; base < 10; base++ {
				fmt.Printf("%d\n", strconv.FormatInt(cas, base))
			}
			fmt.Printf("%d, %d, %d\n", cas, limit, i)
		}

	}

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
	scanner.Split(bufio.ScanWords)
	start := 0

	for scanner.Scan() {
		if start == 0 {
			start = 1
			continue
		}

		N, _ := strconv.Atoi(scanner.Text())
		scanner.Scan()
		J, _ := strconv.Atoi(scanner.Text())

		printCoinJams(N, J)
		//fmt.Printf("%d %d", N, J)
		start++
	}
}
