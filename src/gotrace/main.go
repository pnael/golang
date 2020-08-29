package main

import (
	"fmt"
	"flag"
	"os"
	"net"
)

func destAddr(dest string) ([4]byte, error) {
    destAddr := [4]byte{0, 0, 0, 0}
    addrs, err := net.LookupHost(dest)
    if err != nil {
            return destAddr, err
    }
    addr := addrs[0]

    ipAddr, err := net.ResolveIPAddr("ip", addr)
    if err != nil {
            return destAddr, err
        }
    copy(destAddr[:], ipAddr.IP.To4())
        return destAddr, nil
}

func main() {
	//TTL := flag.Int("TTL",0,"Set the TTL to start with")

	flag.Parse()

	if len(os.Args) != 2 {
		fmt.Println("Expecting one hostname")
		os.Exit(1)
	}

	hostname := os.Args[1]
	RemoteIP,err := net.LookupHost(hostname)
	if err != nil {
		fmt.Println(RemoteIP)
	}

	destIP,_ := destAddr(hostname)
	fmt.Println(destIP)

}

