package main

import (
  "fmt"
  "golang.org/x/net/icmp" 
)
var ListenAddr = "0.0.0.0"

func Ping(addr string) (error) {
  fmt.Printl(addr)
}
func main() {
	fmt.Println("Ping")

	flag1ptr := flag.String("flag1", "No", " First arg")
	flag.Parse()

	fmt.Println("Flag 1", *flag1ptr)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
}
