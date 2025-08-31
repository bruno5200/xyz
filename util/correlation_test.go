package util_test

import (
	"testing"
	"time"

	u "github.com/bruno5200/xyz/util"
)

func TestGenCorrelationId(t *testing.T) {

	// Generate 1000 correlation IDs and ensure they are unique
	correlationIds := make(map[string]bool)

	for i := 0; i < 1000; i++ {

		id := u.GenCorrelationId()

		if len(id) != int(u.CorrelationLength) {
			t.Errorf("Expected length %d, got %d", u.CorrelationLength, len(id))
		}

		if correlationIds[id] {
			t.Errorf("Duplicate correlation ID found: %s", id)
		}

		correlationIds[id] = true
	}

	t.Logf("Generated %d unique correlation IDs", len(correlationIds))

	// Sleep for a short duration to ensure different seeds in fallback method
	time.Sleep(10 * time.Millisecond)
}

func TestAltGenCorrelationId(t *testing.T) {

	// Generate 1000 correlation IDs using the fallback method and ensure they are unique
	correlationIds := make(map[string]bool)
	for i := 0; i < 1000; i++ {
		id := u.AltGenCorrelationId()
		if len(id) != int(u.CorrelationLength) {
			t.Errorf("Expected length %d, got %d", u.CorrelationLength, len(id))
		}
		if correlationIds[id] {
			t.Errorf("Duplicate correlation ID found: %s", id)
		}
		correlationIds[id] = true
	}
	
	t.Logf("Generated %d unique correlation IDs using fallback method", len(correlationIds))

	// Sleep for a short duration to ensure different seeds in fallback method
	time.Sleep(10 * time.Millisecond)
}