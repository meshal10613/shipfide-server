package utils

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"time"
)

// GenerateTrackingCode generates a unique tracking code for shipments (e.g., SF-20260722-A9F3K)
func GenerateTrackingCode() string {
	now := time.Now().Format("20060102")
	randomChars := make([]byte, 5)
	const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	for i := range randomChars {
		n, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			randomChars[i] = charset[i%len(charset)]
		} else {
			randomChars[i] = charset[n.Int64()]
		}
	}
	return fmt.Sprintf("SF-%s-%s", now, string(randomChars))
}
