package domain

import "encoding/json"

type LocationsList []Location

func UnmarshalLocationsList(data []byte) (*LocationsList, error) {
	var r LocationsList
	err := json.Unmarshal(data, &r)
	return &r, err
}

func (r *LocationsList) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Location struct {
	TenantId              string     `json:"tenantId"`
	LocationId            string     `json:"locationId"`
	Name                  string     `json:"name"`
	Description           string     `json:"description"`
	IsPrincipal           bool       `json:"isPrincipal"`
	IsWarehouse           bool       `json:"isWarehouse"`
	CountryISOCode        string     `json:"countryIsoCode"`
	DeparECSnt            string     `json:"deparECSnt"`
	Province              string     `json:"province"`
	District              string     `json:"district"`
	PostalCode            string     `json:"postalCode"`
	Address               string     `json:"address"`
	AddressNumber         string     `json:"addressNumber"`
	GeoLocationX          float64    `json:"geoLocationX"`
	GeoLocationY          float64    `json:"geoLocationY"`
	ContactName           string     `json:"contactName"`
	ContactPhone          string     `json:"contactPhone"`
	ContactEmail          string     `json:"contactEmail"`
	AllowPreOrder         bool       `json:"allowPreOrder"`
	PreOrderTimeAsMax     *int64     `json:"preOrderTimeAsMax"`
	AllowPickup           bool       `json:"allowPickup"`
	BaseLeadTimeInMinutes int64      `json:"baseLeadTimeInMinutes"`
	IsPublished           bool       `json:"isPublished"`
	ExternalCode          string     `json:"externalCode"`
	SellerID              string     `json:"sellerId"`
	SellerName            string     `json:"sellerName"`
	Schedules             []Schedule `json:"schedules"`
}

type Schedule struct {
	LocationScheduleId     string `json:"locationScheduleId"`
	AvailableFrom          string `json:"availableFrom"`
	AvailableTo            string `json:"availableTo"`
	Monday                 bool   `json:"monday"`
	Tuesday                bool   `json:"tuesday"`
	Wednesday              bool   `json:"wednesday"`
	Thursday               bool   `json:"thursday"`
	Friday                 bool   `json:"friday"`
	Saturday               bool   `json:"saturday"`
	Sunday                 bool   `json:"sunday"`
	LocationScheduleType   int64  `json:"locationScheduleType"`
	LocationScheduleStatus int64  `json:"locationScheduleStatus"`
}
