package client

import (
	"crypto/tls"
	"net/http"
	"net/url"
	"sync"
	"time"

	d "github.com/bruno5200/xyz/quickcomm/domain"
	"github.com/google/uuid"
)

var (
	muAuth = &sync.Mutex{}
)

const (
	headerContentType   string = "Content-Type"
	headerAuthorization string = "Authorization"
	headerXOrgId        string = "X-ORG-ID"
	mimeApllicationForm string = "application/x-www-form-urlencoded"
	
	allPath             string = "all"
	authPath            string = "connect/token"
	categoryPath        string = "catalog/categories"
	locationPath        string = "locations"
	orderPath           string = "sales/orders"
	orderGroupPath      string = orderPath + "/ordergroups"
	organizationPath    string = "organizations"
	productPath         string = "catalog/products"
	sellerPath          string = "catalog/sellers"
	tenantPath          string = "tenants"
)

type Quickcomm interface {
	Categories() (*d.Categories, error)
	Category(categoryId uuid.UUID) (*d.Category, error)
	Locations(sellerId uuid.UUID) (*d.LocationsList, error)
	Order(orderId uuid.UUID) (*d.Order, error)
	OrderByName(orderName string) (*d.Order, error)
	Orders() (*d.OrdersList, error)
	Organization() (*d.Organization, error)
	Product(productId uuid.UUID) (*d.Product, error)
	Products() (*d.ProductsList, error)
	Seller(sellerId uuid.UUID) (*d.Seller, error)
	Sellers() (*d.SellersList, error)
	Tenant() (*d.Tenant, error)
}

type client struct {
	url          *url.URL
	client       *http.Client
	clientId     string
	clientSecret string
	xOrgId       string
	accessToken  string
	tokenType    string
	expiration   time.Time
}

// Creates a new Quickcomm client
//
// usage:
//
//	client := client.NewQuickcommClient(serviceBaseURL, clientId, clientSecret, xOrgId)
func NewQuickcommClient(serviceBaseURL, clientId, clientSecret, xOrgId string) (Quickcomm, error) {

	cId, err := uuid.Parse(clientId)

	if err != nil || cId == uuid.Nil {
		return nil, ErrInvalidClientId
	}

	xId, err := uuid.Parse(xOrgId)

	if err != nil || xId == uuid.Nil {
		return nil, ErrInvalidXOrgId
	}

	if clientSecret == "" {
		return nil, ErrInvalidClientSecret
	}

	url, err := url.Parse(serviceBaseURL)

	if err != nil {
		return nil, err
	}

	client := client{url: url, client: &http.Client{Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}}, clientId: clientId, clientSecret: clientSecret, xOrgId: xOrgId}

	go client.authMechanism()

	return &client, nil
}

func (c *client) sendRequest(resquest *http.Request) (*http.Response, error) {
	return c.client.Do(resquest)
}
