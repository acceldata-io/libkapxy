// Acceldata Inc. and its affiliates.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// 	Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
