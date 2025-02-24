package client

import (
	"fmt"
	"io"
	"log"
	"net/http"

	d "github.com/bruno5200/xyz/quickcomm/domain"
	"github.com/google/uuid"
)

// Returns a list of locations
//
// usage:
//
//	client := client.NewQuickcommClient(serviceBaseURL, clientId, clientSecret, xOrgId)
//	locations, err := client.Locations(sellerId)
func (c *client) Locations(sellerId uuid.UUID) (*d.LocationsList, error) {

	c.authMechanism()

	url := c.url.JoinPath(sellerPath).JoinPath(sellerId.String()).JoinPath(locationPath).String()
	log.Printf("URL: %s", url)

	req, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		log.Printf("Error making request: %s", err)
		return nil, err
	}

	req.Header.Set(headerAuthorization, c.tokenType+" "+c.accessToken)

	resp, err := c.sendRequest(req)

	if err != nil {
		log.Printf("Error sending request: %s", err)
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		if body, err := io.ReadAll(resp.Body); err != nil {
			log.Printf("Error location response: %s", body)
		}
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	// log.Printf("Body: %s", body)

	return d.UnmarshalLocationsList(body)
}
