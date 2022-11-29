package libkapxy

import (
	"fmt"

	kafka "github.com/acceldata-io/libkapxy/protocols/kafka"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)

func packetDecode(pktType string, pktPort int, packet gopacket.Packet, produceReqChan chan Kaptured, payloadSize int32, skipPayloadSize bool) {
	if len(packet.Data()) != 0 {

		// Check for errors
		err := packet.ErrorLayer()
		if err == nil {

			// Let's see if the packet is IP
			ipLayer := packet.Layer(layers.LayerTypeIPv4)
			if ipLayer != nil {

				// Let's see if packet is IPv4
				ip, _ := ipLayer.(*layers.IPv4)
				if ip != nil {

					// Let's see if the packet is TCP
					tcpLayer := packet.Layer(layers.LayerTypeTCP)
					if tcpLayer != nil {
						tcp, _ := tcpLayer.(*layers.TCP)

						// Process the Application Layer
						applicationLayer := packet.ApplicationLayer()
						if applicationLayer != nil {

							// Use Kafka Protocol
							kproto := kafka.New(pktType, pktPort)
							msgSize, apiVersion, producerTopics := kproto.Unmarshal(uint16(tcp.DstPort), applicationLayer.Payload(), payloadSize, skipPayloadSize)

							// Extract Kafka Producer Meta
							if producerTopics != nil && len(producerTopics) > 0 {
								// Data from: ip.SrcIP, tcp.SrcPort, ip.DstIP, tcp.DstPort
								// API Version: apiVersion
								// MSG Length: msgSize
								// Producer -> Topic: producerTopics
								pointTags := make(map[string]string)
								pointFields := make(map[string]interface{})
								for producer, topics := range producerTopics {
									for _, topic := range topics {
										pointTags["topic"] = topic
										pointTags["producer"] = producer
										pointFields["producer"] = producer
									}
								}
								pointFields["host"] = fmt.Sprintf("%v", ip.SrcIP)
								pointFields["apiVersion"] = apiVersion
								pointFields["messageSize"] = msgSize
								produceReqChan <- Kaptured{PointTags: pointTags, PointFields: pointFields, TimeStamp: packet.Metadata().Timestamp}
							}
						}
					}
				}
			}
		}
	}
}
