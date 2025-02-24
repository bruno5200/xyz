package domain

import (
	"errors"

	"github.com/google/uuid"
)

var ErrInvalidClientEmail = errors.New("invalid client email") // invalid client email

type Client struct {
	TenantId               uuid.UUID `json:"tenantId"`
	ClientId               uuid.UUID `json:"clientId"`
	ClientSecret           string    `json:"clientSecret,omitempty"`
	GrantType              string    `json:"grantType,omitempty"`
	Email                  string    `json:"email"`
	FirstName              string    `json:"firstName"`
	LastName               string    `json:"lastName"`
	Phone                  *string   `json:"phone,omitempty"`
	IdentificationNumber   string    `json:"identificationNumber"`
	IdentificationType     string    `json:"identificationType"`
	IdentificationTypeName string    `json:"identificationTypeName"`
	EntityName             string    `json:"entityName"`
	CountryISOCode         string    `json:"countryIsoCode"`
	UserID                 uuid.UUID `json:"userId"`
	ClientType             int64     `json:"clientType"`
	ClientTypeName         string    `json:"clientTypeName"`
	ClientStatus           int64     `json:"clientStatus"`
	ClientStatusName       string    `json:"clientStatusName"`
}
