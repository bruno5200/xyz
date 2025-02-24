package domain_test

import (
	"log"
	"os"
	"testing"

	d "github.com/bruno5200/xyz/quickcomm/domain"
)

func TestUnmarshalCategories(t *testing.T) {
	tests := []struct {
		name     string
		JSONpath string
	}{
		{"Categories", "./../../testdata/quickcomm/categories.json"},
	}

	for i := range tests {

		test := tests[i]

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			// size of file
			fileInfo, err := os.Stat(test.JSONpath)

			if err != nil {
				t.Errorf("Error reading file %s", err)
				return
			}

			log.Printf("File size: %d", fileInfo.Size())

			// read data from json file in testdata folder
			jsonFile, err := os.Open(test.JSONpath)

			if err != nil {
				t.Errorf("Error reading file %s", err)
				return
			}
			defer jsonFile.Close()

			data := make([]byte, int(fileInfo.Size()))

			l, err := jsonFile.Read(data)

			if err != nil {
				t.Errorf("Error reading file %s", err)
				return
			}

			log.Printf("Read %d bytes: %q", l, data[:l])

			categories, err := d.UnmarshalCategories(data)

			if err != nil {
				t.Errorf("Error unmarshalling category %s", err)
				return
			}

			for _, category := range *categories {
				log.Printf("CategoryId: %s, Name: %s", category.CategoryId, category.Name)
			}
		})
	}
}

func TestUnmarshallCategory(t *testing.T) {
	tests := []struct {
		name     string
		JSONpath string
	}{
		{"Category", "./../../testdata/category/3cc983c9-fd98-4950-88d4-6eeb46e73eed.json"},
	}

	for i := range tests {

		test := tests[i]

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			// size of file
			fileInfo, err := os.Stat(test.JSONpath)

			if err != nil {
				t.Errorf("Error reading file %s", err)
				return
			}

			log.Printf("File size: %d", fileInfo.Size())

			// read data from json file in testdata folder
			jsonFile, err := os.Open(test.JSONpath)

			if err != nil {
				t.Errorf("Error reading file %s", err)
				return
			}
			defer jsonFile.Close()

			data := make([]byte, int(fileInfo.Size()))

			l, err := jsonFile.Read(data)

			if err != nil {
				t.Errorf("Error reading file %s", err)
				return
			}

			log.Printf("Read %d bytes: %q", l, data[:l])

			category, err := d.UnmarshalCategory(data)

			if err != nil {
				t.Errorf("Error unmarshalling category %s", err)
				return
			}

			log.Printf("CategoryId: %s, Name: %s", category.CategoryId, category.Name)
		})
	}
}

func TestUnmarshalLocationsList(t *testing.T) {
	tests := []struct {
		name     string
		JSONpath string
	}{
		{"Locations", "./../../testdata/quickcomm/locations.json"},
	}

	for i := range tests {

		test := tests[i]

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			// size of file
			fileInfo, err := os.Stat(test.JSONpath)

			if err != nil {
				t.Errorf("Error reading file %s", err)
				return
			}

			log.Printf("File size: %d", fileInfo.Size())

			// read data from json file in testdata folder
			jsonFile, err := os.Open(test.JSONpath)

			if err != nil {
				t.Errorf("Error reading file %s", err)
				return
			}
			defer jsonFile.Close()

			data := make([]byte, int(fileInfo.Size()))

			l, err := jsonFile.Read(data)

			if err != nil {
				t.Errorf("Error reading file %s", err)
				return
			}

			log.Printf("Read %d bytes: %q", l, data[:l])

			locations, err := d.UnmarshalLocationsList(data)

			if err != nil {
				t.Errorf("Error unmarshalling location %s", err)
				return
			}

			for _, location := range *locations {
				log.Printf("LocationId: %s, Name: %s", location.LocationId, location.Name)
			}
		})
	}
}

func TestUnmarshalOrdersList(t *testing.T) {
	tests := []struct {
		name     string
		JSONpath string
	}{
		{"Orders", "./../../testdata/quickcomm/orders.json"},
	}

	for i := range tests {

		test := tests[i]

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			// size of file
			fileInfo, err := os.Stat(test.JSONpath)

			if err != nil {
				t.Errorf("Error reading file %s", err)
				return
			}

			log.Printf("File size: %d", fileInfo.Size())

			// read data from json file in testdata folder
			jsonFile, err := os.Open(test.JSONpath)

			if err != nil {
				t.Errorf("Error reading file %s", err)
				return
			}
			defer jsonFile.Close()

			data := make([]byte, int(fileInfo.Size()))

			l, err := jsonFile.Read(data)

			if err != nil {
				t.Errorf("Error reading file %s", err)
				return
			}

			log.Printf("Read %d bytes: %q", l, data[:l])

			orders, err := d.UnmarshalOrdersList(data)

			if err != nil {
				t.Errorf("Error unmarshalling order %s", err)
				return
			}

			for _, order := range orders.Orders {
				log.Printf("OrderId: %s, ClientId: %s, SellerId: %s", order.SaleOrderId, order.ClientId, order.SellerId)
			}
		})
	}
}

func TestUnmarshalOrder(t *testing.T) {
	tests := []struct {
		name     string
		JSONpath string
	}{
		{"Order", "./../../testdata/order/bf884f52-79be-4ccb-ac2c-2bb1f0441d33.json"},
		{"Order2", "./../../testdata/order/44c17d65-bb58-447a-974a-4bfebe611292.json"},
	}

	for i := range tests {

		test := tests[i]

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			// size of file
			fileInfo, err := os.Stat(test.JSONpath)

			if err != nil {
				t.Errorf("Error reading file %s", err)
				return
			}

			log.Printf("File size: %d", fileInfo.Size())

			// read data from json file in testdata folder
			jsonFile, err := os.Open(test.JSONpath)

			if err != nil {
				t.Errorf("Error reading file %s", err)
				return
			}
			defer jsonFile.Close()

			data := make([]byte, int(fileInfo.Size()))

			l, err := jsonFile.Read(data)

			if err != nil {
				t.Errorf("Error reading file %s", err)
				return
			}

			log.Printf("Read %d bytes: %q", l, data[:l])

			order, err := d.UnmarshalOrder(data)

			if err != nil {
				t.Errorf("Error unmarshalling order %s", err)
				return
			}

			log.Printf("OrderId: %s, ClientId: %s, SellerId: %s", order.SaleOrderId, order.ClientId, order.SellerId)
		})
	}
}

func TestUnmarshalOrganization(t *testing.T) {
	tests := []struct {
		name     string
		JSONpath string
	}{
		{"Organization", "./../../testdata/quickcomm/organization.json"},
	}

	for i := range tests {

		test := tests[i]

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			// size of file
			fileInfo, err := os.Stat(test.JSONpath)

			if err != nil {
				t.Errorf("Error reading file %s", err)
				return
			}

			log.Printf("File size: %d", fileInfo.Size())

			// read data from json file in testdata folder
			jsonFile, err := os.Open(test.JSONpath)

			if err != nil {
				t.Errorf("Error reading file %s", err)
				return
			}
			defer jsonFile.Close()

			data := make([]byte, int(fileInfo.Size()))

			l, err := jsonFile.Read(data)

			if err != nil {
				t.Errorf("Error reading file %s", err)
				return
			}

			log.Printf("Read %d bytes: %q", l, data[:l])

			organization, err := d.UnmarshalOrganization(data)

			if err != nil {
				t.Errorf("Error unmarshalling price %s", err)
				return
			}

			log.Printf("Organization: %s, Name: %s", organization.Organization.OrganizationId, organization.Organization.Name)
		})
	}
}

func TestUnmarshalProductsList(t *testing.T) {
	tests := []struct {
		name     string
		JSONpath string
	}{
		{"Products", "./../../testdata/quickcomm/products.json"},
	}

	for i := range tests {

		test := tests[i]

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			// size of file
			fileInfo, err := os.Stat(test.JSONpath)

			if err != nil {
				t.Errorf("Error reading file %s", err)
				return
			}

			log.Printf("File size: %d", fileInfo.Size())

			// read data from json file in testdata folder
			jsonFile, err := os.Open(test.JSONpath)

			if err != nil {
				t.Errorf("Error reading file %s", err)
				return
			}
			defer jsonFile.Close()

			data := make([]byte, int(fileInfo.Size()))

			l, err := jsonFile.Read(data)

			if err != nil {
				t.Errorf("Error reading file %s", err)
				return
			}

			log.Printf("Read %d bytes: %q", l, data[:l])

			productsList, err := d.UnmarshalProductsList(data)

			if err != nil {
				t.Errorf("Error unmarshalling products %s", err)
				return
			}

			for _, product := range productsList.Products {
				log.Printf("Product: %s, Name: %s, Value: %s", product.ProductId, product.Name, product.BrandName)
				for _, sku := range product.Skus {
					log.Printf("Sku: %s, Name: %s", sku.Sku, sku.Name)
				}
			}
		})
	}
}

func TestUnmarshalProduct(t *testing.T) {
	tests := []struct {
		name     string
		JSONpath string
	}{
		{"Tu Celular seguro II", "./../../testdata/product/3c5cefe8-4afc-4fc7-c9dc-08dd0a3dd19b.json"},
		{"Tu Celular seguro I", "./../../testdata/product/d3385b8e-636e-4e87-540b-08dd0a3dd192.json"},
		{"Tu Celular seguro.", "./../../testdata/product/902ca49f-7df5-4762-0d54-08dd05bd6ec8.json"},
		{"Tu Celular seguro III", "./../../testdata/product/a6e5f300-bda9-4cc1-72fe-08dd0a3dd189.json"},
		{"Tu Celular seguro IV", "./../../testdata/product/5fe57c49-8faf-4a34-9bfb-08dd0a3dd19e.json"},
		{"Tu Deporte Seguro Plan Premium Mensual", "./../../testdata/product/1640c613-4957-47e0-54ff-08dd3bc28a8a.json"},
		{"Crediseguros test", "./../../testdata/product/061f4f3d-df45-44ff-d8ad-08dd0a3e84b2.json"},
		{"Tu Mascota Segura - Plan Premium Anual", "./../../testdata/product/ddcb2778-b382-4536-cd93-08dd3bc48736.json"},
		{"Tu Deporte Seguro - Plan Premium Anual", "./../../testdata/product/289257aa-b264-4781-9d84-08dd3bc31d55.json"},
	}

	for i := range tests {

		test := tests[i]

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			// size of file
			fileInfo, err := os.Stat(test.JSONpath)

			if err != nil {
				t.Errorf("Error reading file %s", err)
				return
			}

			log.Printf("File size: %d", fileInfo.Size())

			// read data from json file in testdata folder
			jsonFile, err := os.Open(test.JSONpath)

			if err != nil {
				t.Errorf("Error reading file %s", err)
				return
			}
			defer jsonFile.Close()

			data := make([]byte, int(fileInfo.Size()))

			l, err := jsonFile.Read(data)

			if err != nil {
				t.Errorf("Error reading file %s", err)
				return
			}

			log.Printf("Read %d bytes: %q", l, data[:l])

			product, err := d.UnmarshalProduct(data)

			if err != nil {
				t.Errorf("Error unmarshalling product %s,%s", test.name, err)
				return
			}

			log.Printf("Product: %s, Value: %s", product.ProductId, product.BrandName)

			for _, sku := range product.Skus {
				log.Printf("Sku: %s, Name: %s, Variant: %s", sku.Sku, sku.Name, sku.VariantName)
			}
		})
	}
}

func TestUnmarshalSellers(t *testing.T) {
	tests := []struct {
		name     string
		JSONpath string
	}{
		{"Sellers", "./../../testdata/quickcomm/sellers.json"},
	}

	for i := range tests {

		test := tests[i]

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			// size of file
			fileInfo, err := os.Stat(test.JSONpath)

			if err != nil {
				t.Errorf("Error reading file %s", err)
				return
			}

			log.Printf("File size: %d", fileInfo.Size())

			// read data from json file in testdata folder
			jsonFile, err := os.Open(test.JSONpath)

			if err != nil {
				t.Errorf("Error reading file %s", err)
				return
			}
			defer jsonFile.Close()

			data := make([]byte, int(fileInfo.Size()))

			l, err := jsonFile.Read(data)

			if err != nil {
				t.Errorf("Error reading file %s", err)
				return
			}

			log.Printf("Read %d bytes: %q", l, data[:l])

			sellers, err := d.UnmarshalSellers(data)

			if err != nil {
				t.Errorf("Error unmarshalling price %s", err)
				return
			}

			for _, seller := range sellers.Sellers {
				log.Printf("Seller: %s, Name: %s", seller.SellerId, seller.Name)
			}
		})
	}
}

func TestUnmarshalSeller(t *testing.T) {
	tests := []struct {
		name     string
		JSONpath string
	}{
		{"Seller", "./../../testdata/seller/c007919e-950e-49ed-98ee-67de60a67299.json"},
	}

	for i := range tests {

		test := tests[i]

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			// size of file
			fileInfo, err := os.Stat(test.JSONpath)

			if err != nil {
				t.Errorf("Error reading file %s", err)
				return
			}

			log.Printf("File size: %d", fileInfo.Size())

			// read data from json file in testdata folder
			jsonFile, err := os.Open(test.JSONpath)

			if err != nil {
				t.Errorf("Error reading file %s", err)
				return
			}
			defer jsonFile.Close()

			data := make([]byte, int(fileInfo.Size()))

			l, err := jsonFile.Read(data)

			if err != nil {
				t.Errorf("Error reading file %s", err)
				return
			}

			log.Printf("Read %d bytes: %q", l, data[:l])

			seller, err := d.UnmarshalSeller(data)

			if err != nil {
				t.Errorf("Error unmarshalling price %s", err)
				return
			}

			log.Printf("Seller: %s, Name: %s", seller.SellerId, seller.Name)
		})
	}
}

func TestUnmarshalTenant(t *testing.T) {
	tests := []struct {
		name     string
		JSONpath string
	}{
		{"Tenant", "./../../testdata/quickcomm/tenant.json"},
	}

	for i := range tests {

		test := tests[i]

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			// size of file
			fileInfo, err := os.Stat(test.JSONpath)

			if err != nil {
				t.Errorf("Error reading file %s", err)
				return
			}

			log.Printf("File size: %d", fileInfo.Size())

			// read data from json file in testdata folder
			jsonFile, err := os.Open(test.JSONpath)

			if err != nil {
				t.Errorf("Error reading file %s", err)
				return
			}
			defer jsonFile.Close()

			data := make([]byte, int(fileInfo.Size()))

			l, err := jsonFile.Read(data)

			if err != nil {
				t.Errorf("Error reading file %s", err)
				return
			}

			log.Printf("Read %d bytes: %q", l, data[:l])

			tenant, err := d.UnmarshalTenant(data)

			if err != nil {
				t.Errorf("Error unmarshalling price %s", err)
				return
			}

			log.Printf("Tenant: %s, Name: %s", tenant.Organization.OrganizationId, tenant.Organization.Name)
		})
	}
}
