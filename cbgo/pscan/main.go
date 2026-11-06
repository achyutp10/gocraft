package main

import (
	"fmt"
	"net"
	"prac/z-backend/cbgo/pscan/port"
	"strconv"
	"time"
)

func scanPort(protocol, hostname string, port int) bool {
	address := hostname + ":" + strconv.Itoa(port)
	conn, err := net.DialTimeout(protocol, address, 60*time.Second)

	if err != nil {
		return false
	}
	defer conn.Close()
	return true
}

func main() {
	fmt.Println("Port Scanning")
	results := port.InitialScan("localhost")
	fmt.Println(results)

	fmt.Println("=================================")

	widescanresults := port.WideScan("localhost")
	fmt.Println(widescanresults)
}
