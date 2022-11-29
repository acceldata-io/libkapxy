package main

import (
	"fmt"
	"log"
	"time"

	"github.com/acceldata-io/libkapxy"

	"github.com/google/gopacket/pcap"
)

func main() {
	//
	fmt.Println("Hello")

	//
	// Configurations
	device := "any"
	snapshotLength := 1024
	skipPayloadSize := false
	promiscuous := false
	filter := "tcp and dst port 6667"
	kafkaPort := 6667
	sniffDirection := "1"
	timeoutSeconds := -1

	// Opens a live packet sniffer handle
	handle, err := pcap.OpenLive(device, int32(snapshotLength), promiscuous, time.Duration(timeoutSeconds)*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()

	// Set Capture Direction
	if sniffDirection == "0" {
		if err := handle.SetDirection(pcap.DirectionInOut); err != nil {
			fmt.Println("ERROR: while setting PCAP Direction. Because: ", err.Error())
			log.Fatal(err)
		}
	} else if sniffDirection == "1" {
		if err := handle.SetDirection(pcap.DirectionIn); err != nil {
			fmt.Println("ERROR: while setting PCAP Direction. Because: ", err.Error())
			log.Fatal(err)
		}
	} else if sniffDirection == "2" {
		if err := handle.SetDirection(pcap.DirectionOut); err != nil {
			fmt.Println("ERROR: while setting PCAP Direction. Because: ", err.Error())
			log.Fatal(err)
		}
	}

	// Set BPF Rule
	if err = handle.SetBPFFilter(filter); err != nil {
		fmt.Println("ERROR: while setting BPF Rule. Because: ", err.Error())
		log.Fatal(err)
	}

	kapCh := make(chan libkapxy.Kaptured)

	kap := libkapxy.Kap{
		Type:            "kafka",
		Port:            kafkaPort,
		Handle:          handle,
		Output:          kapCh,
		SkipPayloadSize: skipPayloadSize,
	}

	// An empty channel to block the process forever not to exit
	done := make(chan bool)

	kap.Run(done)

	fmt.Println("Starting to read packets...")
	go func() {
		for pr := range kapCh {
			fmt.Println(pr.TimeStamp, " ", pr.PointFields, " ", pr.PointTags)
		}
	}()

	<-done // Blocks forever
}
