package client

import "errors"

var (
	ErrInvalidClientId     = errors.New("invalid client id")                     // invalid client id
	ErrInvalidClientSecret = errors.New("invalid client secret")                 // invalid client secret
	ErrInvalidXOrgId       = errors.New("invalid x-org-id")                      // invalid x-org-id
	ErrClientAuth          = errors.New("error authenticating quickcomm client") // error authenticating quickcomm client
	ErrClientOrder         = errors.New("error retrieving order")                // error retrieving order
	ErrClientOrders        = errors.New("error retrieving orders")               // error retrieving orders
	ErrClientLocation      = errors.New("error retrieving location")             // error retrieving location
	ErrClientProduct       = errors.New("error retrieving product")              // error retrieving product
	ErrClientProducts      = errors.New("error retrieving products")             // error retrieving products
	ErrClientSeller        = errors.New("error retrieving seller")               // error retrieving seller
	ErrClientSellers       = errors.New("error retrieving sellers")              // error retrieving sellers
	ErrClientTenant        = errors.New("error retrieving tenant")               // error retrieving tenant
)

// IsInvalidClientIdError returns true if the error is an invalid client id error
func IsClientAuthError(err error) bool {
	return errors.Is(err, ErrClientAuth)
}

// IsInvalidClientSecretError returns true if the error is an invalid client secret error
func IsClientOrderError(err error) bool {
	return errors.Is(err, ErrClientOrder)
}

// IsInvalidXOrgIdError returns true if the error is an invalid x-org-id error
func IsClientOrdersError(err error) bool {
	return errors.Is(err, ErrClientOrders)
}

// IsClientAuthError returns true if the error is an error authenticating quickcomm client
func IsClientLocationError(err error) bool {
	return errors.Is(err, ErrClientLocation)
}

// IsClientOrderError returns true if the error is an error retrieving order
func IsClientProductError(err error) bool {
	return errors.Is(err, ErrClientProduct)
}

// IsClientOrdersError returns true if the error is an error retrieving orders
func IsClientProductsError(err error) bool {
	return errors.Is(err, ErrClientProducts)
}

// IsClientLocationError returns true if the error is an error retrieving location
func IsClientSellerError(err error) bool {
	return errors.Is(err, ErrClientSeller)
}

// IsClientOrderError returns true if the error is an error retrieving seller
func IsClientSellersError(err error) bool {
	return errors.Is(err, ErrClientSellers)
}

// IsClientLocationError returns true if the error is an error retrieving tenant
func IsClientTenantError(err error) bool {
	return errors.Is(err, ErrClientTenant)
}
