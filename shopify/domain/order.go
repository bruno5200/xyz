package domain

import (
	"encoding/json"
	"time"
)

func UnmarshalOrder(data []byte) (*Order, error) {
	var r OrderResponse
	err := json.Unmarshal(data, &r)
	return &r.Order, err
}

func (r *OrderResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type OrderResponse struct {
	Order Order `json:"order"`
}

type Order struct {
	Id                                    int64          `json:"id"`
	AdminGraphqlAPIId                     string         `json:"admin_graphql_api_id"`
	AppId                                 int64          `json:"app_id"`
	BrowserIP                             string         `json:"browser_ip"`
	BuyerAcceptsMarketing                 bool           `json:"buyer_accepts_marketing"`
	CancelReason                          string         `json:"cancel_reason,omitempty"`
	CancelledAt                           time.Time      `json:"cancelled_at,omitempty"`
	CartToken                             string         `json:"cart_token"`
	CheckoutId                            int64          `json:"checkout_id"`
	CheckoutToken                         string         `json:"checkout_token"`
	ClientDetails                         ClientDetails  `json:"client_details"`
	ClosedAt                              time.Time      `json:"closed_at,omitempty"`
	Company                               string         `json:"company,omitempty"`
	ConfirmationNumber                    string         `json:"confirmation_number"`
	Confirmed                             bool           `json:"confirmed"`
	ContactEmail                          string         `json:"contact_email"`
	CreatedAt                             time.Time      `json:"created_at"`
	Currency                              Currency       `json:"currency"`
	CurrentSubtotalPrice                  string         `json:"current_subtotal_price"`
	CurrentSubtotalPriceSet               Set            `json:"current_subtotal_price_set"`
	CurrentTotalDiscounts                 string         `json:"current_total_discounts"`
	CurrentTotalDiscountsSet              Set            `json:"current_total_discounts_set"`
	CurrentTotalPrice                     string         `json:"current_total_price"`
	CurrentTotalPriceSet                  Set            `json:"current_total_price_set"`
	CurrentTotalTax                       string         `json:"current_total_tax"`
	CurrentTotalTaxSet                    Set            `json:"current_total_tax_set"`
	CustomerLocale                        string         `json:"customer_locale"`
	DutiesIncluded                        bool           `json:"duties_included"`
	Email                                 string         `json:"email"`
	EstimatedTaxes                        bool           `json:"estimated_taxes"`
	FinancialStatus                       string         `json:"financial_status"`
	LandingSite                           string         `json:"landing_site"`
	MerchantBusinessEntityId              string         `json:"merchant_business_entity_id"`
	Name                                  string         `json:"name"`
	Number                                int64          `json:"number"`
	OrderNumber                           int64          `json:"order_number"`
	OrderStatusURL                        string         `json:"order_status_url"`
	PaymentGatewayNames                   []string       `json:"payment_gateway_names"`
	Phone                                 string         `json:"phone,omitempty"`
	PresentmentCurrency                   Currency       `json:"presentment_currency"`
	ProcessedAt                           time.Time      `json:"processed_at"`
	ReferringSite                         string         `json:"referring_site"`
	SourceName                            string         `json:"source_name"`
	SubtotalPrice                         string         `json:"subtotal_price"`
	SubtotalPriceSet                      Set            `json:"subtotal_price_set"`
	Tags                                  string         `json:"tags"`
	TaxExempt                             bool           `json:"tax_exempt"`
	TaxesIncluded                         bool           `json:"taxes_included"`
	Test                                  bool           `json:"test"`
	Token                                 string         `json:"token"`
	TotalCashRoundingPaymentAdjustmentSet Set            `json:"total_cash_rounding_payment_adjustment_set"`
	TotalCashRoundingRefundAdjustmentSet  Set            `json:"total_cash_rounding_refund_adjustment_set"`
	TotalDiscounts                        string         `json:"total_discounts"`
	TotalDiscountsSet                     Set            `json:"total_discounts_set"`
	TotalLineItemsPrice                   string         `json:"total_line_items_price"`
	TotalLineItemsPriceSet                Set            `json:"total_line_items_price_set"`
	TotalOutstanding                      string         `json:"total_outstanding"`
	TotalPrice                            string         `json:"total_price"`
	TotalPriceSet                         Set            `json:"total_price_set"`
	TotalShippingPriceSet                 Set            `json:"total_shipping_price_set"`
	TotalTax                              string         `json:"total_tax"`
	TotalTaxSet                           Set            `json:"total_tax_set"`
	TotalTipReceived                      string         `json:"total_tip_received"`
	TotalWeight                           int64          `json:"total_weight"`
	UpdatedAt                             time.Time      `json:"updated_at"`
	BillingAddress                        Address        `json:"billing_address"`
	Customer                              Customer       `json:"customer"`
	LineItems                             []LineItem     `json:"line_items"`
	ShippingAddress                       Address        `json:"shipping_address"`
	ShippingLines                         []ShippingLine `json:"shipping_lines"`
}

type Address struct {
	FirstName   string  `json:"first_name"`
	Address1    string  `json:"address1"`
	Phone       string  `json:"phone"`
	City        string  `json:"city"`
	Country     string  `json:"country"`
	LastName    string  `json:"last_name"`
	Address2    string  `json:"address2,omitempty"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
	Name        string  `json:"name"`
	CountryCode string  `json:"country_code"`
	Id          int64   `json:"id,omitempty"`
	CustomerId  int64   `json:"customer_id,omitempty"`
	CountryName string  `json:"country_name,omitempty"`
	Default     bool    `json:"default,omitempty"`
}

type ClientDetails struct {
	AcceptLanguage string `json:"accept_language"`
	BrowserIP      string `json:"browser_ip"`
	UserAgent      string `json:"user_agent"`
}

type Set struct {
	ShopMoney        Money `json:"shop_money"`
	PresentmentMoney Money `json:"presentment_money"`
}

type Money struct {
	Amount       string   `json:"amount"`
	CurrencyCode Currency `json:"currency_code"`
}

type Customer struct {
	Id                    int64            `json:"id"`
	Email                 string           `json:"email"`
	CreatedAt             time.Time        `json:"created_at"`
	UpdatedAt             time.Time        `json:"updated_at"`
	FirstName             string           `json:"first_name"`
	LastName              string           `json:"last_name"`
	State                 string           `json:"state"`
	Note                  string           `json:"note,omitempty"`
	VerifiedEmail         bool             `json:"verified_email"`
	TaxExempt             bool             `json:"tax_exempt"`
	Phone                 string           `json:"phone,omitempty"`
	EmailMarketingConsent MarketingConsent `json:"email_marketing_consent"`
	SMSMarketingConsent   MarketingConsent `json:"sms_marketing_consent,omitempty"`
	Tags                  string           `json:"tags"`
	Currency              Currency         `json:"currency"`
	AdminGraphqlAPIId     string           `json:"admin_graphql_api_id"`
	DefaultAddress        Address          `json:"default_address"`
}

type MarketingConsent struct {
	State                string    `json:"state"`
	OptInLevel           string    `json:"opt_in_level"`
	ConsentUpdatedAt     time.Time `json:"consent_updated_at,omitempty"`
	ConsentCollectedFrom string    `json:"consent_collected_from,omitempty"`
}

type LineItem struct {
	Id                         int64      `json:"id"`
	AdminGraphqlAPIId          string     `json:"admin_graphql_api_id"`
	CurrentQuantity            int64      `json:"current_quantity"`
	FulfillableQuantity        int64      `json:"fulfillable_quantity"`
	FulfillmentService         string     `json:"fulfillment_service"`
	GiftCard                   bool       `json:"gift_card"`
	Grams                      int64      `json:"grams"`
	Name                       string     `json:"name"`
	Price                      string     `json:"price"`
	PriceSet                   Set        `json:"price_set"`
	ProductExists              bool       `json:"product_exists"`
	ProductId                  int64      `json:"product_id"`
	Properties                 []Property `json:"properties"`
	Quantity                   int64      `json:"quantity"`
	RequiresShipping           bool       `json:"requires_shipping"`
	Sku                        string     `json:"sku"`
	Taxable                    bool       `json:"taxable"`
	Title                      string     `json:"title"`
	TotalDiscount              string     `json:"total_discount"`
	TotalDiscountSet           Set        `json:"total_discount_set"`
	VariantId                  int64      `json:"variant_id"`
	VariantInventoryManagement string     `json:"variant_inventory_management"`
	VariantTitle               string     `json:"variant_title,omitempty"`
	Vendor                     string     `json:"vendor"`
}

type Property struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type ShippingLine struct {
	Id                 int64  `json:"id"`
	CarrierIdentifier  string `json:"carrier_identifier"`
	Code               string `json:"code"`
	DiscountedPrice    string `json:"discounted_price"`
	DiscountedPriceSet Set    `json:"discounted_price_set"`
	IsRemoved          bool   `json:"is_removed"`
	Phone              string `json:"phone,omitempty"`
	Price              string `json:"price"`
	PriceSet           Set    `json:"price_set"`
	Source             string `json:"source"`
	Title              string `json:"title"`
}

type Currency string

const (
	Bob Currency = "BOB"
)
