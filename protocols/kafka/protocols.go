// Package protocols / Kafka is a cut-short version from Sarama module
// https://github.com/Shopify/sarama/blob/main/LICENSE
package protocols

type Protocol interface {
	Unmarshal(dstPort uint16, payload []byte, payloadSize int32, skipPayloadSize bool) (int, int16, map[string][]string)
}

func New(proto string, serverPort int) Protocol {
	switch proto {
	case "kafka":
		return &kafka{serverPort: serverPort}

	default:
		return nil
	}
}
