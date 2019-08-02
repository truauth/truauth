package lib

import (
	"net/http"
	"github.com/truauth/truauth/pkg/utilities"
)

// AccessTokenIntrospectionRequest type used to make introspection request
type AccessTokenIntrospectionRequest struct {
	Token 			string
	ClientSecret 	string 
	ClientID		string
}

// InitAccessTokenIntrospectionRequest creates a resquest from the http input
func InitAccessTokenIntrospectionRequest(req *http.Request) (*AccessTokenIntrospectionRequest, *AuthErrorResponse) {
	tokenRequest := &AccessTokenIntrospectionRequest{
		Token: utilities.GetParam(req, "token"),
		ClientSecret: utilities.GetParam(req, "client_secret"),
		ClientID: utilities.GetParam(req, "client_id"),
	}

	if !tokenRequest.areParamsValid() {
		return nil, CreateError(InvalidRequest, "missing required paramater(s)")
	}

	return tokenRequest, nil
}

func (request *AccessTokenIntrospectionRequest) areParamsValid() bool {
	return request.ClientSecret != "" &&
	request.Token != "" &&
	request.ClientID != ""
}