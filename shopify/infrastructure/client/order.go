package client

import (
	"fmt"
	"io"
	"log"
	"net/http"

	d "github.com/bruno5200/xyz/shopify/domain"
)

func (c *client) Order(orderId int64) (*d.Order, error) {

	req, err := http.NewRequest(http.MethodGet, c.url.JoinPath(orderPath).JoinPath(fmt.Sprintf("%d.%s", orderId, formatJSON)).String(), nil)

	if err != nil {
		log.Printf("Error making request: %s", err)
		return nil, err
	}

	req.Header.Set(headerShopifyAccessToken, c.accessToken)

	resp, err := c.client.Do(req)

	if err != nil {
		log.Printf("Error sending request: %s", err)
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		if body, err := io.ReadAll(resp.Body); err != nil {
			log.Printf("Error order response: %s", body)
		}
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	// log.Printf("Body: %s", body)

	return d.UnmarshalOrder(body)
}
