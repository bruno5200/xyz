package client

import "errors"

var (
	ErrInvalidStoreName   = errors.New("invalid store name")
	ErrInvalidAccessToken = errors.New("invalid access token")
	ErrInvalidApiKey      = errors.New("invalid api key")
	ErrInvalidApiPass     = errors.New("invalid api pass")
	ErrInvalidApiVersion  = errors.New("invalid api version")
)

func IsInvalidStoreNameError(err error) bool {
	return errors.Is(err, ErrInvalidStoreName)
}

func IsInvalidAccessTokenError(err error) bool {
	return errors.Is(err, ErrInvalidAccessToken)
}

func IsInvalidApiKeyError(err error) bool {
	return errors.Is(err, ErrInvalidApiKey)
}

func IsInvalidApiPassError(err error) bool {
	return errors.Is(err, ErrInvalidApiPass)
}

func IsInvalidApiVersionError(err error) bool {
	return errors.Is(err, ErrInvalidApiVersion)
}
