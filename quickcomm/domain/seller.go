package domain

import (
	"encoding/json"

	"github.com/google/uuid"
)

func UnmarshalSellers(data []byte) (*SellersList, error) {
	var r SellersList
	err := json.Unmarshal(data, &r)
	return &r, err
}

func (r *SellersList) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type SellersList struct {
	Sellers        []Seller `json:"results"`
	CurrentPage    int64    `json:"currentPage"`
	PageCount      int64    `json:"pageCount"`
	PageSize       int64    `json:"pageSize"`
	RowCount       int64    `json:"rowCount"`
	Count          int64    `json:"count"`
	FirstRowOnPage int64    `json:"firstRowOnPage"`
	LastRowOnPage  int64    `json:"lastRowOnPage"`
}

func UnmarshalSeller(data []byte) (*Seller, error) {
	var r Seller
	err := json.Unmarshal(data, &r)
	return &r, err
}

func (r *Seller) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Seller struct {
	TenantId                    string     `json:"tenantId,omitempty"`
	RefID                       string     `json:"refId"`
	SellerId                    uuid.UUID  `json:"sellerId"`
	Name                        string     `json:"name"`
	Description                 string     `json:"description"`
	Phone                       *string    `json:"phone,omitempty"`
	TenantProviderId            *string    `json:"tenantProviderId,omitempty"`
	CompanyName                 *string    `json:"companyName,omitempty"`
	CompanyTaxId                *string    `json:"companyTaxId,omitempty"`
	ContactName                 string     `json:"contactName"`
	ContactEmail                string     `json:"contactEmail"`
	ContactPhone                string     `json:"contactPhone"`
	ContactIdentificationNumber string     `json:"contactIdentificationNumber"`
	ExchangesAndReturns         *string    `json:"exchangesAndReturns,omitempty"`
	DeliveryPolicy              *string    `json:"deliveryPolicy,omitempty"`
	PrivacyAndSecurityPolicy    *string    `json:"privacyAndSecurityPolicy,omitempty"`
	WebSite                     *string    `json:"webSite,omitempty"`
	SellerStatus                int64      `json:"sellerStatus"`
	SellerStatusName            string     `json:"sellerStatusName"`
	Active                      bool       `json:"active"`
	SellerType                  int64      `json:"sellerType"`
	SellerTypeName              string     `json:"sellerTypeName"`
	MetaTitle                   string     `json:"metaTitle"`
	MetaDescription             string     `json:"metaDescription"`
	DisableAutoApproveProduct   bool       `json:"disableAutoApproveProduct"`
	CountryISOCode              string     `json:"countryIsoCode"`
	CurrencyISOCode             string     `json:"currencyIsoCode"`
	CommissionType              int64      `json:"commissionType"`
	CommissionTypeName          string     `json:"commissionTypeName"`
	BaseCommission              float64    `json:"baseCommission"`
	BaseMinimumOrderValue       float64    `json:"baseMinimumOrderValue"`
	BaseMaximumOrderValue       *float64   `json:"baseMaximumOrderValue,omitempty"`
	BaseShippingCost            float64    `json:"baseShippingCost"`
	BaseDeliveryTimeInMinutes   int64      `json:"baseDeliveryTimeInMinutes"`
	BaseLeadTimeInMinutes       int64      `json:"baseLeadTimeInMinutes"`
	Logo                        *string    `json:"logo,omitempty"`
	Banner                      *string    `json:"banner,omitempty"`
	Icon                        *string    `json:"icon,omitempty"`
	Slug                        string     `json:"slug"`
	Rating                      int64      `json:"rating"`
	TotalReviews                int64      `json:"totalReviews"`
	CountReviews                int64      `json:"countReviews"`
	ExchangeRate                *float64   `json:"exchangeRate,omitempty"`
	EmailNotification           *string    `json:"emailNotification,omitempty"`
	Facebook                    *string    `json:"facebook,omitempty"`
	Twitter                     *string    `json:"twitter,omitempty"`
	LinkedIn                    *string    `json:"linkedIn,omitempty"`
	Telegram                    *string    `json:"telegram,omitempty"`
	Whatsapp                    *string    `json:"whatsapp,omitempty"`
	Instagram                   *string    `json:"instagram,omitempty"`
	AccountManagerID            *uuid.UUID `json:"accountManagerId,omitempty"`
	AccountManagerEmail         *string    `json:"accountManagerEmail,omitempty"`
	AccountManagerName          *string    `json:"accountManagerName,omitempty"`
	NotifyToAccountManager      bool       `json:"notifyToAccountManager"`
	AllowPreOrder               bool       `json:"allowPreOrder"`
	PreOrderTimeInAdvance       *int64     `json:"preOrderTimeInAdvance,omitempty"`
	PreOrderTimeAsMax           *int64     `json:"preOrderTimeAsMax,omitempty"`
	AllowPickup                 bool       `json:"allowPickup"`
	AllowDelivery               bool       `json:"allowDelivery"`
	AllowPickupAtOtherLocations bool       `json:"allowPickupAtOtherLocations"`
	BankName                    *string    `json:"bankName,omitempty"`
	BankAccountNumber           *string    `json:"bankAccountNumber,omitempty"`
	BankAccountNumber2          *string    `json:"bankAccountNumber2,omitempty"`
	ProductOrServiceName        *string    `json:"productOrServiceName,omitempty"`
	HasDetraction               bool       `json:"hasDetraction"`
	DetractionAccountNumber     *string    `json:"detractionAccountNumber,omitempty"`
	DetractionRate              float64    `json:"detractionRate"`
	PaymentDocumentType         *string    `json:"paymentDocumentType,omitempty"`
	PaymentCurrencyISOCode      *string    `json:"paymentCurrencyIsoCode,omitempty"`
	HasBankAccount              bool       `json:"hasBankAccount"`
	WorkOtherEcommerce          *string    `json:"workOtherEcommerce,omitempty"`
	ExternalEcommerceName       *string    `json:"externalEcommerceName,omitempty"`
	OwnLogistics                bool       `json:"ownLogistics"`
	UseHibryLogistics           bool       `json:"useHibryLogistics"`
	OnlySendOffer               bool       `json:"onlySendOffer"`
}
