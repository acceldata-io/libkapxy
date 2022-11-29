// Package protocols / Kafka is a cut-short version from Sarama module
// https://github.com/Shopify/sarama/blob/main/LICENSE
package protocols

// RequiredAcks is used in Produce Requests to tell the broker how many replica acknowledgements
// it must see before responding. Any of the constants defined here are valid. On broker versions
// prior to 0.8.2.0 any other positive int16 is also valid (the broker will wait for that many
// acknowledgements) but in 0.8.2.0 and later this will raise an exception (it has been replaced
// by setting the `min.isr` value in the brokers configuration).
type RequiredAcks int16

const (
	// NoResponse doesn't send any response, the TCP ACK is all you get.
	NoResponse RequiredAcks = 0
	// WaitForLocal waits for only the local commit to succeed before responding.
	WaitForLocal RequiredAcks = 1
	// WaitForAll waits for all in-sync replicas to commit before responding.
	// The minimum number of in-sync replicas is configured on the broker via
	// the `min.insync.replicas` configuration key.
	WaitForAll RequiredAcks = -1
)

type ProduceRequest struct {
	TransactionalID *string
	RequiredAcks    RequiredAcks
	Timeout         int32
	Version         int16 // v1 requires Kafka 0.9, v2 requires Kafka 0.10, v3 requires Kafka 0.11
	// records         map[string]map[int32]Records
	Topics []string
}

func (r *ProduceRequest) decode(pd packetDecoder, version int16) error {
	r.Version = version

	if version >= 3 {
		id, err := pd.getNullableString()
		if err != nil {
			return err
		}
		r.TransactionalID = id
	}
	requiredAcks, err := pd.getInt16()
	if err != nil {
		return err
	}
	r.RequiredAcks = RequiredAcks(requiredAcks)
	if r.Timeout, err = pd.getInt32(); err != nil {
		return err
	}
	topicCount, err := pd.getArrayLength()
	if err != nil {
		return err
	}
	if topicCount == 0 {
		return nil
	}

	// r.records = make(map[string]map[int32]Records)
	for i := 0; i < topicCount; i++ {
		topic, err := pd.getString()
		if err != nil {
			return err
		}
		r.Topics = append(r.Topics, topic)

		partitionCount, err := pd.getArrayLength()
		if err != nil {
			return err
		}
		// r.records[topic] = make(map[int32]Records)

		for j := 0; j < partitionCount; j++ {
			// partition, err := pd.getInt32()
			_, err := pd.getInt32()
			if err != nil {
				return err
			}
			size, err := pd.getInt32()
			if err != nil {
				return err
			}
			// recordsDecoder, err := pd.getSubset(int(size))
			_, err = pd.getSubset(int(size))
			if err != nil {
				return err
			}
			// var records Records
			// if err := records.decode(recordsDecoder); err != nil {
			// 	return err
			// }
			// r.records[topic][partition] = records
		}
	}

	return nil
}

func (r *ProduceRequest) key() int16 {
	return 0
}

func (r *ProduceRequest) version() int16 {
	return r.Version
}

func (r *ProduceRequest) requiredVersion() KafkaVersion {
	switch r.Version {
	case 1:
		return V0_9_0_0
	case 2:
		return V0_10_0_0
	case 3:
		return V0_11_0_0
	default:
		return MinVersion
	}
}
