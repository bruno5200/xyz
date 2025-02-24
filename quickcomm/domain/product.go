package domain

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

func UnmarshalProductsList(data []byte) (*ProductsList, error) {
	var r ProductsList
	err := json.Unmarshal(data, &r)
	return &r, err
}

func (r *ProductsList) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type ProductsList struct {
	Products       []Product `json:"results"`
	CurrentPage    int64     `json:"currentPage"`
	PageCount      int64     `json:"pageCount"`
	PageSize       int64     `json:"pageSize"`
	RowCount       int64     `json:"rowCount"`
	Count          int64     `json:"count"`
	FirstRowOnPage int64     `json:"firstRowOnPage"`
	LastRowOnPage  int64     `json:"lastRowOnPage"`
}

func UnmarshalProduct(data []byte) (*Product, error) {
	var r Product
	err := json.Unmarshal(data, &r)
	return &r, err
}

func (r *Product) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Product struct {
	TenantId                  uuid.UUID         `json:"tenantId"`
	ProductId                 uuid.UUID         `json:"productId"`
	BrandId                   uuid.UUID         `json:"brandId"`
	BrandName                 string            `json:"brandName"`
	IsOwner                   bool              `json:"isOwner"`
	SellerId                  uuid.UUID         `json:"sellerId"`
	SellerName                string            `json:"sellerName"`
	SellerCommissionType      int64             `json:"sellerCommissionType"`
	SellerCommissionTypeName  string            `json:"sellerCommissionTypeName"`
	ShopName                  string            `json:"shopName,omitempty"`
	ShowInCatalog             bool              `json:"showInCatalog"`
	HasVariations             bool              `json:"hasVariations"`
	Name                      string            `json:"name"`
	Description               string            `json:"description,omitempty"`
	Slug                      string            `json:"slug"`
	BasePrice                 float64           `json:"basePrice"`
	SpecialPrice              float64           `json:"specialPrice,omitempty"`
	DisablePrice              bool              `json:"disablePrice"`
	SmartPrice                bool              `json:"smartPrice,omitempty"`
	BaseCommission            float64           `json:"baseCommission,omitempty"`
	ProductCategories         []ProductCategory `json:"productCategories,omitempty"`
	SellerBaseCommission      float64           `json:"sellerBaseCommission,omitempty"`
	CategoryBaseCommission    float64           `json:"categoryBaseCommission,omitempty"`
	ConditionType             float64           `json:"conditionType"`
	ConditionTypeName         string            `json:"conditionTypeName"`
	ProductType               int64             `json:"productType"`
	ProductTypeName           string            `json:"productTypeName"`
	AllowStorePickup          bool              `json:"allowStorePickup"`
	AllowHomeDelivery         bool              `json:"allowHomeDelivery"`
	AllowSaveAndSubscription  bool              `json:"allowSaveAndSubscription"`
	AllowPurchaseWithoutStock bool              `json:"allowPurchaseWithoutStock"`
	ApplyTaxes                bool              `json:"applyTaxes"`
	AdditionalNoteRequired    bool              `json:"additionalNoteRequired,omitempty"`
	CountryISOCode            string            `json:"countryIsoCode"`
	CurrencyISOCode           string            `json:"currencyIsoCode"`
	MetaTitle                 string            `json:"metaTitle,omitempty"`
	MetaDescription           string            `json:"metaDescription,omitempty"`
	Keywords                  string            `json:"keywords,omitempty"`
	ProductStatus             int64             `json:"productStatus"`
	ProductStatusName         string            `json:"productStatusName"`
	Active                    bool              `json:"active"`
	ExternalCode              string            `json:"externalCode,omitempty"`
	Order                     int64             `json:"order,omitempty"`
	TotalViews                int64             `json:"totalViews,omitempty"`
	TotalLikes                int64             `json:"totalLikes,omitempty"`
	Rating                    int64             `json:"rating,omitempty"`
	TotalReviews              int64             `json:"totalReviews,omitempty"`
	CountReviews              int64             `json:"countReviews,omitempty"`
	DefaultImage              string            `json:"defaultImage"`
	Images                    []Image           `json:"images,omitempty"`
	Skus                      []Skus            `json:"skus,omitempty"`
	CreatedOn                 time.Time         `json:"createdOn"`
	UpdatedOn                 time.Time         `json:"updatedOn,omitempty"`
	CreatedBy                 string            `json:"createdBy"`
	UpdatedBy                 string            `json:"updatedBy,omitempty"`
	ProductOptionType         int64             `json:"productOptionType,omitempty"`
	ProductOptionTypeName     string            `json:"productOptionTypeName,omitempty"`
	ShowInCatalogWithOutStock bool              `json:"showInCatalogWithOutStock,omitempty"`
	CostPrice                 int64             `json:"costPrice,omitempty"`
	TotalStock                int64             `json:"totalStock,omitempty"`
	TotalVariants             int64             `json:"totalVariants,omitempty"`
}

type Image struct {
	Title           string `json:"title"`
	ImageID         string `json:"imageId"`
	IsPrimary       bool   `json:"isPrimary"`
	Order           int64  `json:"order"`
	URLLinkOriginal string `json:"urlLinkOriginal"`
	URLLinkThumb    string `json:"urlLinkThumb"`
	URLLinkCatalog  string `json:"urlLinkCatalog"`
	URLLinkProduct  string `json:"urlLinkProduct"`
	URLLinkZoom     string `json:"urlLinkZoom"`
	HasError        bool   `json:"hasError"`
}

type ProductCategory struct {
	CategoryId   uuid.UUID `json:"categoryId"`
	CategoryName string    `json:"categoryName"`
	ProductId    uuid.UUID `json:"productId"`
}

type Skus struct {
	TenantId                  uuid.UUID   `json:"tenantId"`
	SkuId                     uuid.UUID   `json:"skuId"`
	ProductId                 uuid.UUID   `json:"productId"`
	VariantName               string      `json:"variantName"`
	Sku                       string      `json:"sku"`
	BarCode                   string      `json:"barCode"`
	Name                      string      `json:"name"`
	SellerId                  uuid.UUID   `json:"sellerId"`
	SellerName                string      `json:"sellerName"`
	CountryISOCode            string      `json:"countryIsoCode"`
	CurrencyISOCode           string      `json:"currencyIsoCode"`
	TrackingStock             bool        `json:"trackingStock"`
	AllowPurchaseWithoutStock bool        `json:"allowPurchaseWithoutStock"`
	Stock                     int64       `json:"stock"`
	MinStock                  int64       `json:"minStock"`
	BasePrice                 float64     `json:"basePrice"`
	SpecialPrice              float64     `json:"specialPrice"`
	LoyaltyValue              float64     `json:"loyaltyValue"`
	EntityStatus              int64       `json:"entityStatus"`
	SkuStatus                 int64       `json:"skuStatus"`
	SkuStatusName             string      `json:"skuStatusName"`
	ExternalCode              string      `json:"externalCode"`
	Order                     int64       `json:"order"`
	Height                    float64     `json:"height"`
	Length                    float64     `json:"length"`
	Width                     float64     `json:"width"`
	Weight                    float64     `json:"weight"`
	UnitMutilplier            int64       `json:"unitMutilplier"`
	MinimumPurchase           int64       `json:"minimumPurchase"`
	PackageHeight             float64     `json:"packageHeight"`
	PackageLength             float64     `json:"packageLength"`
	PackageWeight             float64     `json:"packageWeight"`
	PackageWidth              float64     `json:"packageWidth"`
	ExcludeShippingCost       bool        `json:"excludeShippingCost"`
	Inventories               []Inventory `json:"inventories"`
	Active                    bool        `json:"active"`
}

type Inventory struct {
	SkuId          uuid.UUID `json:"skuId"`
	LocationId     uuid.UUID `json:"locationId"`
	LocationName   string    `json:"locationName"`
	Stock          int64     `json:"stock"`
	StockReserved  int64     `json:"stockReserved"`
	StockAvailable int64     `json:"stockAvailable"`
	GeoLocationY   float64   `json:"geoLocationY"`
	GeoLocationX   float64   `json:"geoLocationX"`
	AllowPickup    bool      `json:"allowPickup"`
	IsPrincipal    bool      `json:"isPrincipal"`
}
