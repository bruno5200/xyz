package client

import (
	"fmt"
	"io"
	"log"
	"net/http"

	d "github.com/bruno5200/xyz/quickcomm/domain"
)

// Retrieves the tenant information
//
// usage:
//
//	client := client.NewQuickcommClient(serviceBaseURL, clientId, clientSecret, xOrgId)
//	tenant, err := client.Tenant()
func (c *client) Tenant() (*d.Tenant, error) {

	c.authMechanism()

	url := c.url.JoinPath(tenantPath).JoinPath(c.xOrgId).String()
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
			log.Printf("Error tenant response: %s", body)
		}
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	// log.Printf("Body: %s", body)

	return d.UnmarshalTenant(body)
}
