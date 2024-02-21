package helper

import (
	"crypto/rand"
	"fmt"
	"hash/crc64"
)

func GenerateRandomBytes(n int) ([]byte, error) {
	if n == 0 {
		n = 8
	}
	randomBytes := make([]byte, n)
	if _, err := rand.Read(randomBytes); err != nil {
		return nil, err
	}
	return randomBytes, nil
}

func ComputeCRC64(data []byte) string {
	table := crc64.MakeTable(crc64.ECMA)
	crcSum := crc64.Checksum(data, table)
	return fmt.Sprintf("%x", crcSum)
}
