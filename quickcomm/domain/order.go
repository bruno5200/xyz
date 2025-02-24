package domain

import (
	"encoding/json"

	"github.com/bruno5200/xyz/util"

	"github.com/google/uuid"
)

func UnmarshalOrdersList(data []byte) (*OrdersList, error) {
	var r OrdersList
	err := json.Unmarshal(data, &r)
	return &r, err
}

func (r *OrdersList) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type OrdersList struct {
	Orders         []Order `json:"results"`
	CurrentPage    int64   `json:"currentPage"`
	PageCount      int64   `json:"pageCount"`
	PageSize       int64   `json:"pageSize"`
	RowCount       int64   `json:"rowCount"`
	Count          int64   `json:"count"`
	FirstRowOnPage int64   `json:"firstRowOnPage"`
	LastRowOnPage  int64   `json:"lastRowOnPage"`
}

type Orders []Order

func UnmarshalOrders(data []byte) (*Orders, error) {
	var r Orders
	err := json.Unmarshal(data, &r)
	return &r, err
}

func (r *Orders) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalOrder(data []byte) (*Order, error) {
	var r Order
	err := json.Unmarshal(data, &r)
	return &r, err
}

func (r *Order) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Order struct {
	TenantId                     uuid.UUID            `json:"tenantId,omitempty"`
	SaleOrderId                  uuid.UUID            `json:"saleOrderId,omitempty"`
	CheckOutId                   uuid.UUID            `json:"checkOutId,omitempty"`
	Source                       string               `json:"source,omitempty"`
	MarketplaceProviderId        uuid.UUID            `json:"marketplaceProviderId,omitempty"`
	MarketplaceProviderTenantId  uuid.UUID            `json:"marketplaceProviderTenantId,omitempty"`
	SaleOrderStatus              int64                `json:"saleOrderStatus,omitempty"`
	SaleOrderStatusName          string               `json:"saleOrderStatusName,omitempty"`
	ShippingOrderType            string               `json:"shippingOrderType,omitempty"`
	SaleOrderExternalId          string               `json:"saleOrderExternalId,omitempty"`
	OrderType                    int64                `json:"orderType,omitempty"`
	EntityName                   string               `json:"entityName,omitempty"`
	EntityIdentificationNumber   string               `json:"entityIdentificationNumber,omitempty"`
	OrderDate                    util.CustomTime      `json:"orderDate,omitempty"`
	ConfirmedOrderDate           util.CustomTime      `json:"confirmedOrderDate,omitempty"`
	ShippingOrderDate            util.CustomTime      `json:"shippingOrderDate,omitempty"`
	ShippingLate                 bool                 `json:"shippingLate,omitempty"`
	EstimatedShippingDate        util.CustomTime      `json:"estimatedShippingDate,omitempty"`
	OrderGroup                   string               `json:"orderGroup,omitempty"`
	OrderNumber                  string               `json:"orderNumber,omitempty"`
	OrderReferenceNumber         string               `json:"orderReferenceNumber,omitempty"`
	SellerId                     uuid.UUID            `json:"sellerId,omitempty"`
	SellerName                   string               `json:"sellerName,omitempty"`
	SellerCompanyName            string               `json:"sellerCompanyName,omitempty"`
	SellerTaxId                  string               `json:"sellerTaxId,omitempty"`
	ClientId                     uuid.UUID            `json:"clientId,omitempty"`
	Client                       Client               `json:"client,omitempty"`
	CountryISOCode               string               `json:"countryIsoCode,omitempty"`
	CurrencyISOCode              string               `json:"currencyIsoCode,omitempty"`
	SubTotal                     float64              `json:"subTotal,omitempty"`
	Shipping                     float64              `json:"shipping,omitempty"`
	ShippingAssumedByMarketplace float64              `json:"shippingAssumedByMarketplace,omitempty"`
	Tax                          float64              `json:"tax,omitempty"`
	Discount                     float64              `json:"discount,omitempty"`
	ServiceFee                   float64              `json:"serviceFee,omitempty"`
	Tips                         float64              `json:"tips,omitempty"`
	Total                        float64              `json:"total,omitempty"`
	TotalUsd                     float64              `json:"totalUsd,omitempty"`
	ExchangeRate                 float64              `json:"exchangeRate,omitempty"`
	IsPrePaid                    bool                 `json:"isPrePaid,omitempty"`
	IsPreOrder                   bool                 `json:"isPreOrder,omitempty"`
	TotalLoyalty                 float64              `json:"totalLoyalty,omitempty"`
	MileUsed                     float64              `json:"mileUsed,omitempty"`
	IsPickUp                     bool                 `json:"isPickUp,omitempty"`
	IsSubscription               bool                 `json:"isSubscription,omitempty"`
	ShippingFirstName            string               `json:"shippingFirstName,omitempty"`
	ShippingLastName             string               `json:"shippingLastName,omitempty"`
	ShippingIdentificationType   string               `json:"shippingIdentificationType,omitempty"`
	ShippingIdentificationNumber string               `json:"shippingIdentificationNumber,omitempty"`
	ShippingAddressLine          string               `json:"shippingAddressLine,omitempty"`
	ShippingReference            string               `json:"shippingReference,omitempty"`
	ShippingCountryISOCode       string               `json:"shippingCountryIsoCode,omitempty"`
	ShippingLatitude             string               `json:"shippingLatitude,omitempty"`
	ShippingLongitude            string               `json:"shippingLongitude,omitempty"`
	ShippingDeparxyznt           string               `json:"shippingDeparxyznt,omitempty"`
	AdditionalProperties         []AdditionalProperty `json:"additionalProperties,omitempty"`
	SaleOrderItems               []SaleOrderItem      `json:"saleOrderItems,omitempty"`
	Trackings                    []Tracking           `json:"trackings,omitempty"`
	Transactions                 []Transaction        `json:"transactions,omitempty"`
	Shipments                    json.RawMessage      `json:"shipments"`
	Seller                       SellerOrder          `json:"seller,omitempty"`
}

type AdditionalProperty struct {
	Group string `json:"group"`
	Key   string `json:"key"`
	Value string `json:"value"`
}

type SaleOrderItem struct {
	SaleOrderItemId         uuid.UUID       `json:"saleOrderItemId"`
	SaleOrderId             uuid.UUID       `json:"saleOrderId"`
	SaleOrderItemStatus     int64           `json:"saleOrderItemStatus"`
	SaleOrderItemStatusName string          `json:"saleOrderItemStatusName"`
	ProductId               uuid.UUID       `json:"productId"`
	ProductExternalId       string          `json:"productExternalId"`
	ProductName             string          `json:"productName"`
	ProductImage            string          `json:"productImage"`
	Weight                  float64         `json:"weight"`
	SubTotal                float64         `json:"subTotal"`
	Discount                float64         `json:"discount"`
	Total                   float64         `json:"total"`
	TotalLoyalty            float64         `json:"totalLoyalty"`
	MileUsed                float64         `json:"mileUsed"`
	AdditionalNote          string          `json:"additionalNote"`
	ShippingCost            float64         `json:"shippingCost"`
	ShippingCostListPrice   float64         `json:"shippingCostListPrice"`
	ShippingDate            util.CustomTime `json:"shippingDate"`
	ShippingDateUntil       util.CustomTime `json:"shippingDateUntil"`
	QuantityToBeDelivered   int64           `json:"quantityToBeDelivered"`
	Category                *string         `json:"category"`
	CategoryIds             *string         `json:"categoryIds"`
	IsPickupStore           bool            `json:"isPickupStore"`
	ShippingType            *string         `json:"shippingType"`
	StorePickupAddressLine  *string         `json:"storePickupAddressLine"`
	StorePickupDistrict     *string         `json:"storePickupDistrict"`
	StorePickupProvince     *string         `json:"storePickupProvince"`
	StorePickupDeparxyznt   *string         `json:"storePickupDeparxyznt"`
	StorePickupPostalCode   *string         `json:"storePickupPostalCode"`
	StorePickupLatitude     float64         `json:"storePickupLatitude"`
	StorePickupLongitude    float64         `json:"storePickupLongitude"`
	ShippingMethodName      string          `json:"shippingMethodName"`
	SellerId                string          `json:"sellerId"`
	SellerName              string          `json:"sellerName"`
	SellerTaxId             string          `json:"sellerTaxId"`
	Variants                []Variant       `json:"variants"`
	Trackings               []Tracking      `json:"trackings"`
	Seller                  SellerOrder     `json:"seller"`
}

type SellerOrder struct {
	SellerId     uuid.UUID   `json:"sellerId"`
	Name         *string     `json:"name"`
	CompanyName  *string     `json:"companyName"`
	ContactName  *string     `json:"contactName"`
	ContactEmail *string     `json:"contactEmail"`
	ContactPhone *string     `json:"contactPhone"`
	Attributes   []Attribute `json:"attributes"`
}

type Attribute struct {
	Id    string `json:"id"`
	Key   string `json:"key"`
	Name  string `json:"name"`
	Value string `json:"value"`
}

type Tracking struct {
	SaleOrderItemTrackingId *string `json:"saleOrderItemTrackingId,omitempty"`
	Title                   string  `json:"title"`
	Content                 *string `json:"content"`
	TrackingType            int64   `json:"trackingType"`
	TrackingTypeName        string  `json:"trackingTypeName"`
	CreatedOn               string  `json:"createdOn"`
	CreatedBy               *string `json:"createdBy"`
	UpdatedEmail            *string `json:"updatedEmail"`
	Completed               bool    `json:"completed"`
	SaleOrderTrackingId     *string `json:"saleOrderTrackingId,omitempty"`
	SaleOrderId             *string `json:"saleOrderId,omitempty"`
}

type Variant struct {
	SaleOrderItemVariantId string    `json:"saleOrderItemVariantId"`
	SaleOrderItemId        string    `json:"saleOrderItemId"`
	SellerId               string    `json:"sellerId"`
	SellerName             string    `json:"sellerName"`
	SkuExternalId          string    `json:"skuExternalId"`
	SkuId                  string    `json:"skuId"`
	Sku                    string    `json:"sku"`
	Name                   string    `json:"name"`
	SkuImage               *string   `json:"skuImage"`
	Quantity               int64     `json:"quantity"`
	QuantityToBeDelivered  int64     `json:"quantityToBeDelivered"`
	QuantityToReturned     int64     `json:"quantityToReturned"`
	BasePrice              float64   `json:"basePrice"`
	SpecialPrice           float64   `json:"specialPrice"`
	FinalPrice             float64   `json:"finalPrice"`
	Height                 float64   `json:"height"`
	Length                 float64   `json:"length"`
	Weight                 float64   `json:"weight"`
	Width                  float64   `json:"width"`
	SubTotal               float64   `json:"subTotal"`
	Discount               float64   `json:"discount"`
	ShippingCost           float64   `json:"shippingCost"`
	ShippingCostListPrice  float64   `json:"shippingCostListPrice"`
	Total                  float64   `json:"total"`
	LoyaltyValue           float64   `json:"loyaltyValue"`
	MileUsedValue          float64   `json:"mileUsedValue"`
	LocationId             uuid.UUID `json:"locationId"`
	LocationName           *string   `json:"locationName"`
}

type Transaction struct {
	SaleOrderTransactionId       string  `json:"saleOrderTransactionId"`
	SaleOrderId                  string  `json:"saleOrderId"`
	TransactionId                string  `json:"transactionId"`
	SaleOrderTransactionType     int64   `json:"saleOrderTransactionType"`
	SaleOrderTransactionTypeName string  `json:"saleOrderTransactionTypeName"`
	ProviderId                   string  `json:"providerId"`
	PaymenxyzthodGroupId         string  `json:"paymenxyzthodGroupId"`
	PaymenxyzthodGroupName       *string `json:"paymenxyzthodGroupName"`
	PaymenxyzthodId              string  `json:"paymenxyzthodId"`
	PaymenxyzthodName            *string `json:"paymenxyzthodName"`
	TransactionDate              string  `json:"transactionDate"`
	Amount                       float64 `json:"amount"`
	UsedMiles                    float64 `json:"usedMiles"`
	PlaceHolder                  *string `json:"placeHolder"`
	CardNumber                   *string `json:"cardNumber"`
	Installments                 int64   `json:"installments"`
	TransactionOperationNumber   string  `json:"transactionOperationNumber"`
	CreatedOn                    string  `json:"createdOn"`
	UpdatedOn                    string  `json:"updatedOn"`
}
