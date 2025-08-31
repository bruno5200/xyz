package util

import (
	crand "crypto/rand"
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
		return AltGenCorrelationId()
	}

	return string(b)
}

// AltGenCorrelationId is a fallback method to generate a correlation ID using math/rand.
//
// This method is less secure and should only be used if crypto/rand fails.
func AltGenCorrelationId() string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, CorrelationLength)
	rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
