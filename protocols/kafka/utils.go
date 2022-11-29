// Package protocols / Kafka is a cut-short version from Sarama module
// https://github.com/Shopify/sarama/blob/main/LICENSE
package protocols

import (
	"errors"
	"math/big"
)

// InBetween checks if the given value is between the range or not
func InBetween(i, min, max int) bool {
	if (i >= min) && (i <= max) {
		return true
	}
	return false
}

// TODO: Optimize the this function, may not be efficient. Just a working code.
func getLen(transIDLenPayload []byte) (int, error) {
	bigIntData := new(big.Int).SetBytes(transIDLenPayload)
	// If the first byte starts with '1' it's a negative number
	// SEE: https://stackoverflow.com/questions/53829120/convert-negative-digit-in-byte-slice-to-int
	// Hexadecimal value of '0x80' is equal to '128' in decimal AND '10000000' in binary
	if len(transIDLenPayload) > 0 && transIDLenPayload[0]&0x80 != 0 {
		tempData := make([]byte, len(transIDLenPayload)+1)
		tempData[0] = 1
		tempTransIDLen := new(big.Int).SetBytes(tempData)
		bigIntData.Sub(tempTransIDLen, bigIntData)

		// Apply negative sign and convert to int
		return int(bigIntData.Neg(bigIntData).Int64()), nil
	}
	return 0, errors.New("not a negative number")
}
