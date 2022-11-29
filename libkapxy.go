// Package libkapxy provides network packet decoder for various protocols
package libkapxy

import (
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

type Kap struct {
	Type            string
	Port            int
	Handle          *pcap.Handle
	Output          chan Kaptured
	SkipPayloadSize bool
}

type Kaptured struct {
	PointTags   map[string]string
	PointFields map[string]interface{}
	TimeStamp   time.Time
}

// Run starts the packet capture
func (k *Kap) Run(done chan bool) {
	go k.pkt(done)
}

func (k *Kap) pkt(done chan bool) {
	// Create a packet source and loop over the packetsource channel
	packetSource := gopacket.NewPacketSource(k.Handle, k.Handle.LinkType())
	for packet := range packetSource.Packets() {
		packetDecode(k.Type, k.Port, packet, k.Output, int32(k.Handle.SnapLen()), k.SkipPayloadSize)
	}

	done <- true
}
