package domain_test

import (
	"log"
	"os"
	"testing"

	"github.com/bruno5200/xyz/shopify/domain"
)

func TestUnmarshalOrder(t *testing.T) {
	tests := []struct {
		name     string
		JSONpath string
	}{
		{"Order1", "./../../testdata/shopify/order/5616862068921.json"},
		{"Order2", "./../../testdata/shopify/order/5617017782457.json"},
	}

	for i := range tests {

		test := tests[i]

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			fileInfo, err := os.Stat(test.JSONpath)

			if err != nil {
				t.Errorf("Error reading file %s", err)
				return
			}

			log.Printf("File size: %d", fileInfo.Size())

			// read data from json file in testdata folder
			jsonFile, err := os.Open(test.JSONpath)

			if err != nil {
				t.Errorf("Error reading file %s", err)
				return
			}
			defer jsonFile.Close()

			data := make([]byte, int(fileInfo.Size()))

			l, err := jsonFile.Read(data)

			if err != nil {
				t.Errorf("Error reading file %s", err)
				return
			}

			log.Printf("Read %d bytes: %q", l, data[:l])

			order, err := domain.UnmarshalOrder(data)

			if err != nil {
				t.Errorf("Error unmarshalling order %s", err)
				return
			}

			log.Printf("OrderId: %d, Number: %d", order.Id, order.OrderNumber)
		})
	}
}
