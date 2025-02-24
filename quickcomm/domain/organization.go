package domain

import (
	"encoding/json"

	"github.com/bruno5200/xyz/util"
)

func UnmarshalOrganization(data []byte) (*Organization, error) {
	var r Organization
	err := json.Unmarshal(data, &r)
	return &r, err
}

func (r *Organization) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Organization struct {
	Organization OrganizationClass `json:"organization"`
	Client       Client            `json:"client"`
}

type OrganizationClass struct {
	OrganizationId                   string          `json:"organizationId"`
	Name                             string          `json:"name"`
	Host                             string          `json:"host"`
	AccountOwnerId                   string          `json:"accountOwnerId"`
	AccountName                      string          `json:"accountName"`
	AccountEmail                     string          `json:"accountEmail"`
	TimeZone                         string          `json:"timeZone"`
	OffSet                           int64           `json:"offSet"`
	CurrencyISO                      string          `json:"currencyISO"`
	CountryISO                       string          `json:"countryISO"`
	CurrencySymbol                   string          `json:"currencySymbol"`
	IsOwner                          bool            `json:"isOwner"`
	TrialExpiredOn                   string          `json:"trialExpiredOn"`
	HostType                         int64           `json:"hostType"`
	Plan                             int64           `json:"plan"`
	Type                             int64           `json:"type"`
	MarketPlaceType                  int64           `json:"marketPlaceType"`
	PlatformType                     int64           `json:"platformType"`
	Price                            float64         `json:"price"`
	CreatedOn                        util.CustomTime `json:"createdOn"`
	UpdatedOn                        util.CustomTime `json:"updatedOn"`
	EntityStatus                     int64           `json:"entityStatus"`
	Logo                             string          `json:"logo"`
	SetupCompleted                   bool            `json:"setupCompleted"`
	Settings                         []Setting       `json:"settings"`
	Clients                          []int64         `json:"clients"`
	HasPaymentAvailable              bool            `json:"hasPaymentAvailable"`
	HasPaymentProviderAvailable      bool            `json:"hasPaymentProviderAvailable"`
	HasLogisticAvailable             bool            `json:"hasLogisticAvailable"`
	HasCMSAvailable                  bool            `json:"hasCMSAvailable"`
	HasNotificationAvailable         bool            `json:"hasNotificationAvailable"`
	HasNotificationProviderAvailable bool            `json:"hasNotificationProviderAvailable"`
	HasMarketplaceAvailable          bool            `json:"hasMarketplaceAvailable"`
	HasCategoriesAvailable           bool            `json:"hasCategoriesAvailable"`
	HasBrandsAvailable               bool            `json:"hasBrandsAvailable"`
	HasProductsAvailable             bool            `json:"hasProductsAvailable"`
}

type Setting struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
