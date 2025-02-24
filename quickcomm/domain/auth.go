package domain

import "encoding/json"

func UnmarshalAuthResponse(data []byte) (*AuthResponse, error) {
	var r AuthResponse
	err := json.Unmarshal(data, &r)
	return &r, err
}

func (r *AuthResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type AuthResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
	TokenType   string `json:"token_type"`
	Scope       string `json:"scope"`
}
