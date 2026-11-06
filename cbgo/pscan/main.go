package main

import (
	"fmt"
	"net"
	"prac/z-backend/cbgo/pscan/port"
	"strconv"
	"time"
)

func ScanPort(protocol, hostname string, port int) bool {
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
	for _, result := range results {
		if result.State == "Open" {
			fmt.Println(result.Port, result.State)
		}
	}

	fmt.Println("=================================")

	widescanresults := port.WideScan("localhost")
	for _, result := range widescanresults {
		if result.State == "Open" {
			fmt.Println(result.Port, result.State)
		}
	}
}
