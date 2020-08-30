package main

import (
	"fmt"
	"flag"
	"os"
	"net"
	"syscall"
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

	destIP,_ := destAddr(hostname)
	sendSocket, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_DGRAM, syscall.IPPROTO_UDP)
	if err != nil {
		fmt.Println("sendSocket error :", err)
		os.Exit(1)
	}
	recvSocket, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_RAW, syscall.IPPROTO_ICMP)
	if err != nil {
		fmt.Println("recvSocket error :", err)
		os.Exit(1)
	}
	defer syscall.Close(recvSocket)
	defer syscall.Close(sendSocket)

	fmt.Println(destIP)

}

