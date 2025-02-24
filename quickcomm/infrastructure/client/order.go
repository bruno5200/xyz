package client

import (
	"fmt"
	"io"
	"log"
	"net/http"

	d "github.com/bruno5200/xyz/quickcomm/domain"
	"github.com/google/uuid"
)

// Returns a list of orders
//
// usage:
//
//	client := client.NewQuickcommClient(serviceBaseURL, clientId, clientSecret, xOrgId)
//	orders, err := client.Orders()
func (c *client) Orders() (*d.OrdersList, error) {

	c.authMechanism()

	url := c.url.JoinPath(orderPath).String()
	log.Printf("URL: %s", url)

	req, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		log.Printf("Error making request: %s", err)
		return nil, err
	}

	req.Header.Set(headerAuthorization, c.tokenType+" "+c.accessToken)
	req.Header.Set(headerXOrgId, c.xOrgId)

	resp, err := c.sendRequest(req)

	if err != nil {
		log.Printf("Error sending request: %s", err)
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		if body, err := io.ReadAll(resp.Body); err != nil {
			log.Printf("Error orders response: %s", body)
		}
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	// log.Printf("Body: %s", body)

	return d.UnmarshalOrdersList(body)
}

// Returns an order by id
//
// usage:
//
//	client := client.NewQuickcommClient(serviceBaseURL, clientId, clientSecret, xOrgId)
//	order, err := client.Order(orderId)
func (c *client) Order(orderId uuid.UUID) (*d.Order, error) {

	c.authMechanism()

	url := c.url.JoinPath(orderPath).JoinPath(orderId.String()).String()
	log.Printf("URL: %s", url)

	req, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		log.Printf("Error making request: %s", err)
		return nil, err
	}

	req.Header.Set(headerAuthorization, c.tokenType+" "+c.accessToken)
	req.Header.Set(headerXOrgId, c.xOrgId)

	resp, err := c.sendRequest(req)

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

// Returns an order by name
//
// usage:
//
//	client := client.NewQuickcommClient(serviceBaseURL, clientId, clientSecret, xOrgId)
//	order, err := client.OrderByName(orderName)
func (c *client) OrderByName(orderName string) (*d.Order, error) {

	c.authMechanism()

	url := c.url.JoinPath(orderGroupPath).JoinPath(orderName).String()
	log.Printf("URL: %s", url)

	req, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		log.Printf("Error making request: %s", err)
		return nil, err
	}

	req.Header.Set(headerAuthorization, c.tokenType+" "+c.accessToken)
	req.Header.Set(headerXOrgId, c.xOrgId)

	resp, err := c.sendRequest(req)

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
	orders, err := d.UnmarshalOrders(body)

	if len(*orders) == 0 {
		return nil, fmt.Errorf("order not found")
	}

	return &(*orders)[0], err
}
