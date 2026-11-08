package main

import (
	"bytes"
	"fmt"
	"log"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

var DevName = "Ethernet"
var Found = false

func main() {
	// 1. find all network interfaces for packet capturing using pcap.FindAllDevs()
	devices, err := pcap.FindAllDevs()
	if err != nil {
		// log.Panic("Unable to fetch network interfacse")
		log.Fatalf("Unable to fetch network interfaces: %v", err)

	}

	for _, ifDev := range devices {
		if ifDev.Name == DevName {
			Found = true
		}
	}

	// 2. If the desired device is found cotinue otherwise exit the program

	if !Found {
		log.Panic("Desired deice not found")
	}

	// 3. Open live capture handle on that network interface using pcap.OpenLive()

	handle, err := pcap.OpenLive(DevName, 1600, false, pcap.BlockForever)
	if err != nil {
		fmt.Println(err)
		log.Panic("Unable to open handle on the device")
	}
	defer handle.Close()

	// 4. Apply BPF (Berkely Packet Filter) on this new handle

	if err := handle.SetBPFFilter("tcp and port 21"); err != nil {
		log.Panic(err)
	}

	// 5. Display all the filtered packets received on channel returned from gopacket.NewPacketSource()

	source := gopacket.NewPacketSource(handle, handle.LinkType())

	for packet := range source.Packets() {
		appLayer := packet.ApplicationLayer()
		if appLayer == nil {
			continue
		}
		data := appLayer.Payload()
		if bytes.Contains(data, []byte("USER")) || bytes.Contains(data, []byte("PASS")) {
			fmt.Println(string(data))
		}

		// fmt.Println(packet)
	}

}

// func main() {
// 	devices, err := pcap.FindAllDevs()

// 	if err != nil {
// 		log.Fatalf("Unable to fetch network interfaces: %v", err)
// 	}

// 	fmt.Println("Available Network Interfaces:")

// 	for i, device := range devices {
// 		fmt.Printf("\nInterface %d\n", i+1)
// 		fmt.Println("Name:", device.Name)
// 		fmt.Println("Description:", device.Description)

// 		for _, address := range device.Addresses {
// 			fmt.Println("Address:", address.IP)
// 		}
// 	}
// }
