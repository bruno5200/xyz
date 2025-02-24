package client

import (
	"fmt"
	"io"
	"log"
	"net/http"

	d "github.com/bruno5200/xyz/quickcomm/domain"
	"github.com/google/uuid"
)

// Returns a list of products
//
// usage:
//
//	client := client.NewQuickcommClient(serviceBaseURL, clientId, clientSecret, xOrgId)
//	products, err := client.Products()
func (c *client) Products() (*d.ProductsList, error) {

	c.authMechanism()

	url := c.url.JoinPath(productPath).String()
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
			log.Printf("Error products response: %s", body)
		}
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	// log.Printf("Body: %s", body)

	return d.UnmarshalProductsList(body)
}

// Returns a product
//
// usage:
//
//	client := client.NewQuickcommClient(serviceBaseURL, clientId, clientSecret, xOrgId)
//	product, err := client.Product(productId)
func (c *client) Product(productId uuid.UUID) (*d.Product, error) {

	c.authMechanism()

	url := c.url.JoinPath(productPath).JoinPath(productId.String()).String()
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
			log.Printf("Error product response: %s", body)
		}
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	// log.Printf("Body: %s", body)

	return d.UnmarshalProduct(body)
}
