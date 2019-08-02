package lib

import (
	"github.com/truauth/truauth/pkg/utilities"
	"net/http"
)

// AuthTokenRequest request made when requesting an auth token
type AuthTokenRequest struct {
	GrantType		string 
	Code 			string 
	RedirectURI 	string 
	ClientID 		string
	ClientSecret 	string
	RefreshToken 	string // required field for the token refresh grant
}

// InitAuthTokenRequest initializes the auth token struct
func InitAuthTokenRequest(req *http.Request) (*AuthTokenRequest, *AuthErrorResponse) {
	authToken := &AuthTokenRequest{
		GrantType: utilities.GetParam(req, "grant_type"),
		Code: utilities.GetParam(req, "code"),
		RedirectURI: utilities.GetParam(req, "redirect_uri"),
		ClientID: utilities.GetParam(req, "client_id"),
		ClientSecret: utilities.GetParam(req, "client_secret"),
		RefreshToken: utilities.GetParam(req, "refresh_token"),
	}

	if !authToken.areParamsValid() {
		return nil, CreateError(InvalidRequest, "missing required paramater")
	}

	return authToken, nil
}

// Validate validates the token request
func (authToken *AuthTokenRequest) Validate(
	authCode *CodeToken,
) *AuthErrorResponse {
	if !authToken.areParamsValid() {
		return CreateError(InvalidRequest, "missing required paramater(s)")
	}

	if !authToken.isSignatureValid(authCode) {
		return CreateError(UnauthorizedClient, "invalid token signature")
	}

	return nil
}

func (authToken *AuthTokenRequest) areParamsValid() bool {
	return authToken.GrantType != "" &&
	authToken.Code != "" &&
	authToken.RedirectURI != "" &&
	authToken.ClientID != "" &&
	authToken.ClientSecret != ""
}

func (authToken *AuthTokenRequest) isSignatureValid(authCode *CodeToken) bool {
//strings.Compare(authToken.RedirectURI , authCode.RedirectURI) == 0 &&

	return !authCode.HasExpired()
}