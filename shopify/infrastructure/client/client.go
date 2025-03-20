package client

import (
	"crypto/tls"
	"net/http"
	"net/url"

	"github.com/bruno5200/xyz/shopify/domain"
)

const (
	protocol                 string = "https"
	formatJSON               string = "json"
	shopifyDomain            string = "myshopify.com"
	apiPath                  string = "admin/api"
	headerShopifyAccessToken string = "X-Shopify-Access-Token"
	orderPath                string = "orders"
)

type Shopify interface {
	Order(orderId int64) (*domain.Order, error)
}

type client struct {
	url    *url.URL
	client *http.Client

	accessToken string
	apiKey      string
	apiPass     string
	apiversion  string
}

func NewShopifyClient(storeName, accessToken, apiKey, apiPass, apiversion string) (Shopify, error) {

	if storeName == "" {
		return nil, ErrInvalidStoreName
	}

	if accessToken == "" {
		return nil, ErrInvalidAccessToken
	}

	if apiKey == "" {
		return nil, ErrInvalidApiKey
	}

	if apiPass == "" {
		return nil, ErrInvalidApiPass
	}

	if apiversion == "" {
		return nil, ErrInvalidApiVersion
	}

	url, err := url.Parse(protocol + "://" + storeName + "." + shopifyDomain)

	if err != nil {
		return nil, err
	}

	url = url.JoinPath(apiPath).JoinPath(apiversion)

	return &client{
		url:         url,
		client:      &http.Client{Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}},
		accessToken: accessToken,
		apiKey:      apiKey,
		apiPass:     apiPass,
		apiversion:  apiversion,
	}, nil
}
