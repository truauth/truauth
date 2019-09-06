package lib

import (
	"github.com/truauth/truauth/pkg/utilities"
)

// AccessTokenResponse access token used to grant resource access
type AccessTokenResponse struct {
	AccessToken  string    `json:"access_token"`
	TokenType    TokenType `json:"token_type"`
	ExpiresIn    int       `json:"expires_in"`
	RefreshToken string    `json:"refresh_token"`
	Scope        string    `json:"scope"`
	ID           string    `json:"id_token"`
}

// CreateAccessTokenResponse creates an access token response
func CreateAccessTokenResponse(
	signedToken string,
	tokenType TokenType,
	expiresIn int,
	refreshToken string,
) *AccessTokenResponse {
	return &AccessTokenResponse{ // todo: support rest of stuff
		AccessToken:  signedToken,
		TokenType:    tokenType,
		ExpiresIn:    expiresIn,
		RefreshToken: refreshToken,
		ID:           utilities.UUID(),
	}
}
