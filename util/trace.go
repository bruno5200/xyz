package util

import "github.com/google/uuid"

// GenTraceId generates a new UUID string to be used as a trace ID.
func GenTraceId() string {
	return uuid.New().String()
}