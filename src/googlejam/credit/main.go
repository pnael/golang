package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)


func findArticles(a int, b []int) (int,int) {

		var c,d int
		for i , price := range b {
			if price > a {
				continue
			}
			sli := b[i+1:]
			for j,t := range sli {
				if price + t == a {
					c = i+1
					d = j+i+2
					continue
				}
			}
		}
		return c,d
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
	//Skip the number of case
  scanner.Scan()
	i:=0
	for scanner.Scan() {

		totalcredit,_ := strconv.Atoi(scanner.Text())
		//Don't need the number of article
		scanner.Scan()

		//Get all the prices
		scanner.Scan()
		articles_text := scanner.Text()

		articles := strings.Split(articles_text, " ")


		var prices []int
		for _, article  := range(articles) {
			art,_ := strconv.Atoi(article)
			prices = append(prices, art)
		}

		a,b := findArticles(totalcredit, prices)
		i++
		fmt.Printf("Case #%d: %d %d\n",i,a,b)

		}
}
