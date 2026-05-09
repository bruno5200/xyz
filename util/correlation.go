package util

import (
	crand "crypto/rand"
	"encoding/hex"
	"math/rand"
	"time"
)

const (
	CorrelationLength uint = 16
)

// GenCorrelationId generates a random correlation ID of length CorrelationLength and using crypto/rand for better randomness.
//
// If crypto/rand fails, it falls back to a less secure method called AltGenCorrelationId which uses math/rand.
func GenCorrelationId() string {
	b := make([]byte, CorrelationLength)

	if _, err := crand.Read(b); err != nil {
		// Fallback to less secure method if crypto/rand fails.
		return AltGenCorrelationId()
	}

	// Encode the random bytes to a hexadecimal string to ensure it's printable.
	// This will result in a string of length CorrelationLength * 2.
	return hex.EncodeToString(b)
}

// The random source is initialized once to ensure better randomness across calls.
var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

// AltGenCorrelationId is a fallback method to generate a correlation ID using math/rand.
//
// This method is less secure and should only be used if crypto/rand fails.
func AltGenCorrelationId() string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	// To maintain a consistent output length with GenCorrelationId (which is hex encoded),
	// we generate a string of length CorrelationLength * 2.
	result := make([]byte, CorrelationLength)
	for i := range result {
		result[i] = letters[seededRand.Intn(len(letters))]
	}
	return string(result)
}
