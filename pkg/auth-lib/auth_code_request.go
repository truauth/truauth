package lib

import (
	"net/http"
	"github.com/truauth/truauth/pkg/utilities"
	"strings"
)

// AuthCodeRequest request made when requesting an auth code
type AuthCodeRequest struct {
	ResponseType 	string
	ClientID		string
	RedirectURI		string
	Scope			[]string
	State			string
}

// InitAuthCode creates an auth code provided http paramaters
func InitAuthCode(req *http.Request) (*AuthCodeRequest, *AuthErrorResponse) {
	scope := utilities.GetParam(req, "scope")

	authCode := &AuthCodeRequest{
		ResponseType: utilities.GetParam(req, "response_type"),
		ClientID:     utilities.GetParam(req, "client_id"),
		Scope:        strings.Split(scope, ","),
		RedirectURI:  utilities.GetParam(req, "redirect_uri"),
		State:        utilities.GetParam(req, "state"),
	}

	if (!authCode.areParamsValid()) {
		return nil, CreateError(InvalidRequest, "missing required paramater(s)")
	}

	return authCode, nil
}

// Validate validates the auth code request
func (request *AuthCodeRequest) Validate() *AuthErrorResponse {
	if !request.areParamsValid() {
		return CreateError(InvalidRequest, "missing required parameter(s)")
	}

	if !request.isOpenIDScopePresent() {
		return CreateError(InvalidRequest, "openid scope is missing")
	}

	return nil
}

func (request *AuthCodeRequest) isOpenIDScopePresent() bool {
	for _, scope := range request.Scope {
		if scope == "openid" {
			return true
		}
	}
	return false
}

func (request *AuthCodeRequest) areParamsValid() bool {
	return request.ClientID != "" &&
	request.RedirectURI != "" &&
	request.ResponseType != "" &&
	len(request.Scope) > 0 &&
	request.State != ""
}