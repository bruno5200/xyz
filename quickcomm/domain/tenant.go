package domain

import (
	"encoding/json"

	"github.com/google/uuid"
)

func UnmarshalTenant(data []byte) (*Tenant, error) {
	var r Tenant
	err := json.Unmarshal(data, &r)
	return &r, err
}

func (r *Tenant) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Tenant struct {
	Organization OrganizationInfo `json:"organization"`
}

type OrganizationInfo struct {
	OrganizationId  uuid.UUID `json:"organizationId"`
	Name            string    `json:"name"`
	Host            string    `json:"host"`
	AccountOwnerId  uuid.UUID `json:"accountOwnerId"`
	AccountName     string    `json:"accountName"`
	AccountEmail    string    `json:"accountEmail"`
	Type            int64     `json:"type"`
	PlatformType    int64     `json:"platformType"`
	MarketPlaceType int64     `json:"marketPlaceType"`
	Logo            string    `json:"logo"`
	ContactPhone    *string   `json:"contactPhone,omitempty"`
	ContactEmail    *string   `json:"contactEmail,omitempty"`
	Settings        []Setting `json:"settings"`
	Clients         []Client  `json:"clients"`
}
