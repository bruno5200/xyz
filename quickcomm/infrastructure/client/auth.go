package client

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"time"

	d "github.com/bruno5200/xyz/quickcomm/domain"
)

func (c *client) auth() (*d.AuthResponse, error) {

	data := url.Values{}
	data.Set("grant_type", "client_credentials")
	data.Set("scope", "api IdentityServerApi")
	data.Set("client_id", c.clientId)
	data.Set("client_secret", c.clientSecret)
	data.Set("x-Org-Id", c.xOrgId)

	payload := bytes.NewBufferString(data.Encode())

	// log.Printf("Payload: %v", payload)

	req, err := http.NewRequest(http.MethodPost, c.url.JoinPath(authPath).String(), payload)
	if err != nil {
		log.Printf("Error making request: %s", err)
		return nil, err
	}

	req.Header.Add(headerContentType, mimeApllicationForm)

	resp, err := c.sendRequest(req)

	if err != nil {
		log.Printf("Error sending request: %s", err)
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {

		if body, err := io.ReadAll(resp.Body); err != nil {
			log.Printf("Error auth response: %s", body)
		}

		return nil, fmt.Errorf("unexpected status code %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	// log.Printf("Body: %s", body)

	return d.UnmarshalAuthResponse(body)
}

func (c *client) authMechanism() {
	muAuth.Lock()
	defer muAuth.Unlock()
	if c.accessToken == "" || time.Now().After(c.expiration) {
		authResponse, err := c.auth()

		if err != nil {
			log.Printf("Error authenticating Quickcomm client: %s", err)
			return
		}

		c.accessToken = authResponse.AccessToken
		c.tokenType = authResponse.TokenType
		c.expiration = time.Now().Add(time.Duration(authResponse.ExpiresIn) * time.Second)

		// log.Printf("Access Token: %s, Expires in: %d, Scope: %s", c.accessToken, authResponse.ExpiresIn, authResponse.Scope)

		log.Println("Quickcomm Client authentication successful")
	}
}
