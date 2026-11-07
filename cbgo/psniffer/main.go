package main

import (
	"fmt"
	"log"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

var DevName = "wlp3s0"
var Found = false

func main() {
	// 1. find all network interfaces for packet capturing using pcap.FindAllDevs()
	devices, err := pcap.FindAllDevs()
	if err != nil {
		log.Panic("Unable to fetch network interfacse")
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

	if err := handle.SetBPFFilter("tcp and port 433"); err != nil {
		log.Panic(err)
	}

	// 5. Display all the filtered packets received on channel returned from gopacket.NewPacketSource()

	source := gopacket.NewPacketSource(handle, handle.LinkType())

	for packets := range source.Packets() {
		fmt.Println(packets)
	}

}
