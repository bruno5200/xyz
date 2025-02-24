package domain

import (
	"encoding/json"

	"github.com/google/uuid"
)

type Categories []Category

func UnmarshalCategories(data []byte) (*Categories, error) {
	var r Categories
	err := json.Unmarshal(data, &r)
	return &r, err
}

func (r *Categories) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalCategory(data []byte) (*Category, error) {
	var r Category
	err := json.Unmarshal(data, &r)
	return &r, err
}

func (r *Category) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Category struct {
	TenantId             uuid.UUID          `json:"tenantId"`
	CategoryId           uuid.UUID          `json:"categoryId"`
	CategoryParentId     uuid.UUID          `json:"categoryParentId"`
	SellerId             uuid.UUID          `json:"sellerId"`
	Name                 string             `json:"name"`
	TreeName             string             `json:"treeName"`
	Description          string             `json:"description"`
	Image                string             `json:"image"`
	Icon                 string             `json:"icon"`
	Slug                 string             `json:"slug"`
	Order                int64              `json:"order"`
	BaseCommission       float64            `json:"baseCommission"`
	CategoryStatus       int64              `json:"categoryStatus"`
	IsFromSync           bool               `json:"isFromSync"`
	CategoryStatusName   CategoryStatusName `json:"categoryStatusName"`
	CategoryCode         string             `json:"categoryCode"`
	CategoryManager      string             `json:"categoryManager"`
	CategoryManagerId    uuid.UUID          `json:"categoryManagerId"`
	CategoryManagerEmail string             `json:"categoryManagerEmail"`
	MetaTitle            string             `json:"metaTitle"`
	MetaDescription      string             `json:"metaDescription"`
	ShowInMenu           bool               `json:"showInMenu"`
	ExternalCode         string             `json:"externalCode"`
	InternalCode         int64              `json:"internalCode"`
}

type CategoryStatusName string

const (
	Active CategoryStatusName = "Active"
)
