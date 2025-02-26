package client_test

import (
	"encoding/base64"
	"log"
	"testing"

	c "github.com/bruno5200/xyz/shopify/infrastructure/client"
)

const (
	storeName   string = "dmall-bolivia"
	accessToken string = "c2hwYXRfZmJlM2JkZDRhM2RlYTczYjIzZGVhNzJmYWZhOTMwNjE="
	apiKey      string = "e8b8ee01be4e56e14d03c1f16604a378"
	apiPass     string = "20d71650763270637ecb1e4ca842e56c"
	apiVersion  string = "2025-01"
)

func TestOrder(t *testing.T) {
	//Table Driven Test
	tests := []struct {
		name    string
		orderId int64
	}{
		{"OneItem", 5617017782457},
		{"TwoItem", 5616862068921},
		// {"Order 3", 3},
	}

	token, err := base64.StdEncoding.DecodeString(accessToken)

	if err != nil {
		t.Errorf("Error: %s", err)
		return
	}

	cl, err := c.NewShopifyClient(storeName, string(token), apiKey, apiPass, apiVersion)

	if err != nil {
		t.Errorf("Error: %s", err)
		return
	}

	for i := range tests {
		test := tests[i]
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			order, err := cl.Order(test.orderId)

			if err != nil {
				t.Errorf("Error: %s", err)
				return
			}

			log.Printf("Order: %d, OrderNumber: %d", order.Id, order.Number)

			for i := range order.LineItems {
				log.Printf("LineItem: %d, Title: %s", order.LineItems[i].Id, order.LineItems[i].Title)
			}
		})
	}
}
