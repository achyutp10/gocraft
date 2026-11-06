package port

import (
	"net"
	"strconv"
	"sync"
	"time"
)

type ScanResult struct {
	Port    string
	State   string
	Service string
}

func ScanPort(protocol, hostname string, port int) ScanResult {
	result := ScanResult{
		Port: strconv.Itoa(port) + "/" + protocol,
	}
	address := hostname + ":" + strconv.Itoa(port)
	conn, err := net.DialTimeout(
		protocol,
		address,
		3*time.Second,
	)

	if err != nil {
		result.State = "Closed"
		return result
	}
	defer conn.Close()
	result.State = "Open"
	return result
}

func InitialScan(hostname string) []ScanResult {
	var wg sync.WaitGroup

	results := make(chan ScanResult, 2048)

	for i := 0; i < 1024; i++ {
		for _, protocol := range []string{"udp", "tcp"} {
			wg.Add(1)

			go func(protocol string, port int) {
				defer wg.Done()
				result := ScanPort(protocol, hostname, port)
				results <- result
			}(protocol, i)
		}
	}
	go func() {
		wg.Wait()
		close(results)
	}()

	var scanResults []ScanResult

	for result := range results {
		scanResults = append(scanResults, result)
	}

	return scanResults
}

func WideScan(hostname string) []ScanResult {
	var wg sync.WaitGroup

	results := make(chan ScanResult, 2048)

	for i := 0; i < 49152; i++ {
		for _, protocol := range []string{"udp", "tcp"} {
			wg.Add(1)

			go func(protocol string, port int) {
				defer wg.Done()
				result := ScanPort(protocol, hostname, port)
				results <- result
			}(protocol, i)
		}
	}
	go func() {
		wg.Wait()
		close(results)
	}()

	var scanResults []ScanResult

	for result := range results {
		scanResults = append(scanResults, result)
	}

	return scanResults
}
