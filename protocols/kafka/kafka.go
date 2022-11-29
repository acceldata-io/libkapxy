// Package protocols / Kafka is a cut-short version from Sarama module
// https://github.com/Shopify/sarama/blob/main/LICENSE
// Kafka Protocol Implementation
// SEE: https://kafka.apache.org/protocol
package protocols

import (
	"bytes"
)

type kafka struct {
	serverPort int
}

func (k *kafka) Unmarshal(dstPort uint16, payload []byte, payloadSize int32, skipPayloadSize bool) (int, int16, map[string][]string) {
	if dstPort == uint16(k.serverPort) {
		return unmarshalRequest(payload, payloadSize, skipPayloadSize)
	}

	return int(0), int16(0), nil
}

func unmarshalRequest(payload []byte, payloadSize int32, skipPayloadSize bool) (int, int16, map[string][]string) {
	// request header
	//===============
	// 0-3 message_size
	// 4-5 api_key
	// 6-7 api_version
	// 8-11 correlation_id
	// 12-13 client_id length

	producerReq := make(map[string][]string)

	payloadReader := bytes.NewReader(payload)

	requestBody, bytesRead, err := decodeRequest(payloadReader, payloadSize, skipPayloadSize)
	if err != nil {
		return int(0), int16(0), nil
	}

	pReq, ok := requestBody.body.(*ProduceRequest)
	if ok != true {
		// This means the interface value doesn't hold the type ProduceRequest
		return int(0), int16(0), nil
	}

	producerReq[requestBody.clientID] = pReq.Topics
	return bytesRead, requestBody.body.version(), producerReq
}
