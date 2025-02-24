package client_test

import (
	"fmt"
	"log"
	"sync"
	"testing"

	c "github.com/bruno5200/xyz/quickcomm/infrastructure/client"
	"github.com/google/uuid"
)

const (
	sellerURL string = "https://sandbox-sellercentral.elgeniox.com"
	clientId  string = "c007919e-950e-49ed-98ee-67de60a67299"
	secret    string = "RmbkTl9yWY0DrpuqWRVnbrjTyA2xGavwjEEpjLbP+Xo="
	orgId     string = "16d172b0-c5df-4597-8daf-eb4461297e35"
)

func TestCategories(t *testing.T) {
	cl, err := c.NewQuickcommClient(sellerURL, clientId, secret, orgId)

	if err != nil {
		t.Errorf("Error: %s", err)
		return
	}

	categories, err := cl.Categories()

	if err != nil {
		t.Errorf("Error: %s", err)
		return
	}

	for _, category := range *categories {
		log.Printf("Category Id: %s, %s", category.CategoryId, category.Name)
	}
}

func TestCategory(t *testing.T) {
	cl, err := c.NewQuickcommClient(sellerURL, clientId, secret, orgId)

	if err != nil {
		t.Errorf("Error: %s", err)
		return
	}

	category, err := cl.Category(uuid.MustParse("2ec9ce91-63a7-4545-b30a-96f34eaf9116"))

	if err != nil {
		t.Errorf("Error: %s", err)
		return
	}

	log.Printf("Category: %s", category.CategoryId)
}

func TestOrders(t *testing.T) {
	cl, err := c.NewQuickcommClient(sellerURL, clientId, secret, orgId)

	if err != nil {
		t.Errorf("Error: %s", err)
		return
	}

	ordersList, err := cl.Orders()

	if err != nil {
		t.Errorf("Error: %s", err)
		return
	}

	for _, order := range ordersList.Orders {
		fmt.Printf("OrderId: %s, Client: %s, OrderGroup: %s\n", order.SaleOrderId, order.ClientId, order.OrderGroup)

		for _, item := range order.SaleOrderItems {
			fmt.Printf("ProductId: %s, ProductName: %s, Location: %s, Quantity: %d\n", item.ProductId, item.ProductName, item.Variants[0].LocationId, item.Variants[0].Quantity)
		}
	}
}

func TestOrder(t *testing.T) {
	cl, err := c.NewQuickcommClient(sellerURL, clientId, secret, orgId)

	if err != nil {
		t.Errorf("Error: %s", err)
		return
	}

	order, err := cl.Order(uuid.MustParse("bf884f52-79be-4ccb-ac2c-2bb1f0441d33"))

	if err != nil {
		t.Errorf("Error: %s", err)
		return
	}

	fmt.Printf("OrderId: %s, Client: %s, OrderGroup: %s\n", order.SaleOrderId, order.ClientId, order.OrderGroup)

	for _, item := range order.SaleOrderItems {
		fmt.Printf("ProductId: %s, ProductName: %s, Location: %s, Quantity: %d\n", item.ProductId, item.ProductName, item.Variants[0].LocationId, item.Variants[0].Quantity)
	}
}

func TestOrderByGroup(t *testing.T) {
	cl, err := c.NewQuickcommClient(sellerURL, clientId, secret, orgId)

	if err != nil {
		t.Errorf("Error: %s", err)
		return
	}

	order, err := cl.OrderByName("4038")

	if err != nil {
		t.Errorf("Error: %s", err)
		return
	}

	fmt.Printf("OrderId: %s, Client: %s, OrderGroup: %s\n", order.SaleOrderId, order.ClientId, order.OrderGroup)

	for _, item := range order.SaleOrderItems {
		fmt.Printf("ProductId: %s, ProductName: %s, Location: %s, Quantity: %d\n", item.ProductId, item.ProductName, item.Variants[0].LocationId, item.Variants[0].Quantity)
	}
}

func TestProducts(t *testing.T) {
	cl, err := c.NewQuickcommClient(sellerURL, clientId, secret, orgId)

	if err != nil {
		t.Errorf("Error: %s", err)
		return
	}

	productsList, err := cl.Products()

	if err != nil {
		t.Errorf("Error: %s", err)
		return
	}

	for _, p := range productsList.Products {
		log.Printf("Id: %s, Name: %s, Seller: %s", p.ProductId, p.Name, p.SellerId)
	}
}

func TestProduct(t *testing.T) {
	cl, err := c.NewQuickcommClient(sellerURL, clientId, secret, orgId)

	if err != nil {
		t.Errorf("Error: %s", err)
		return
	}

	product, err := cl.Product(uuid.MustParse("902ca49f-7df5-4762-0d54-08dd05bd6ec8"))

	if err != nil {
		t.Errorf("Error: %s", err)
		return
	}

	log.Printf("Id: %s, Name: %s, Seller: %s", product.ProductId, product.Name, product.SellerId)
	for _, sku := range product.Skus {
		log.Printf("Sku: %s, Name: %s", sku.Sku, sku.Name)
	}
}

func TestOrganization(t *testing.T) {
	cl, err := c.NewQuickcommClient(sellerURL, clientId, secret, orgId)

	if err != nil {
		t.Errorf("Error: %s", err)
		return
	}

	organization, err := cl.Organization()

	if err != nil {
		t.Errorf("Error: %s", err)
		return
	}

	log.Printf("Organization: %s", organization.Organization.OrganizationId)
}

func TestSellers(t *testing.T) {
	cl, err := c.NewQuickcommClient(sellerURL, clientId, secret, orgId)

	if err != nil {
		t.Errorf("Error: %s", err)
		return
	}

	sellersList, err := cl.Sellers()

	if err != nil {
		t.Errorf("Error: %s", err)
		return
	}

	for _, seller := range sellersList.Sellers {
		log.Printf("Seller: %s, Name: %s", seller.SellerId, seller.Name)
	}
}

func TestTenant(t *testing.T) {
	cl, err := c.NewQuickcommClient(sellerURL, clientId, secret, orgId)

	if err != nil {
		t.Errorf("Error: %s", err)
		return
	}

	tenant, err := cl.Tenant()

	if err != nil {
		t.Errorf("Error: %s", err)
		return
	}

	log.Printf("Tenant: %s", tenant.Organization.OrganizationId)
}

func TestSeller(t *testing.T) {
	cl, err := c.NewQuickcommClient(sellerURL, clientId, secret, orgId)

	if err != nil {
		t.Errorf("Error: %s", err)
		return
	}

	seller, err := cl.Seller(uuid.MustParse("c007919e-950e-49ed-98ee-67de60a67299"))

	if err != nil {
		t.Errorf("Error: %s", err)
		return
	}

	log.Printf("Seller: %s, Name: %s", seller.SellerId, seller.Name)
}

func TestIntegrationSellerCentral(t *testing.T) {
	cl, err := c.NewQuickcommClient(sellerURL, clientId, secret, orgId)

	if err != nil {
		t.Errorf("Error: %s", err)
		return
	}

	sellersList, err := cl.Sellers()

	if err != nil {
		t.Errorf("Error: %s", err)
		return
	}

	wgSeller := sync.WaitGroup{}

	for _, seller := range sellersList.Sellers {

		wgSeller.Add(1)

		go func(id uuid.UUID) {
			defer wgSeller.Done()

			seller, err := cl.Seller(id)

			if err != nil {
				t.Errorf("Error: %s", err)
				return
			}

			fmt.Printf("SellerId: %s, Name: %s\n", seller.SellerId, seller.Name)

		}(seller.SellerId)
	}

	wgLocation := sync.WaitGroup{}

	for _, seller := range sellersList.Sellers {

		wgLocation.Add(1)

		go func(id uuid.UUID) {
			defer wgLocation.Done()

			locations, err := cl.Locations(id)

			if err != nil {
				t.Errorf("Error: %s", err)
				return
			}

			for _, location := range *locations {
				fmt.Printf("LocationId: %s, Name: %s\n", location.LocationId, location.Name)
			}

		}(seller.SellerId)
	}

	ordersList, err := cl.Orders()

	if err != nil {
		t.Errorf("Error: %s", err)
		return
	}

	wgOrder := sync.WaitGroup{}

	for _, order := range ordersList.Orders {

		wgOrder.Add(1)

		go func(id uuid.UUID) {
			defer wgOrder.Done()

			order, err := cl.Order(id)

			if err != nil {
				t.Errorf("Error: %s", err)
				return
			}

			fmt.Printf("OrderId: %s, Client: %s, OrderGroup: %s\n", order.SaleOrderId, order.ClientId, order.OrderGroup)

		}(order.SaleOrderId)
	}

	productsList, err := cl.Products()

	if err != nil {
		t.Errorf("Error: %s", err)
		return
	}

	wgProducts := sync.WaitGroup{}

	for _, p := range productsList.Products {

		wgProducts.Add(1)

		go func(id uuid.UUID) {
			defer wgProducts.Done()

			product, err := cl.Product(id)

			if err != nil {
				t.Errorf("Error: %s", err)
				return
			}

			fmt.Printf("ProductId: %s, Name: %s, Seller: %s\n", product.ProductId, product.Name, product.SellerId)

		}(p.ProductId)
	}

	tenant, err := cl.Tenant()

	if err != nil {
		t.Errorf("Error: %s", err)
		return
	}

	fmt.Printf("TenantId: %s\n", tenant.Organization.OrganizationId)

	organization, err := cl.Organization()

	if err != nil {
		t.Errorf("Error: %s", err)
		return
	}

	fmt.Printf("Organization: %s\n", organization.Organization.OrganizationId)

	wgSeller.Wait()
	wgLocation.Wait()
	wgOrder.Wait()
	wgProducts.Wait()
}
