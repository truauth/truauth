package main

import (
	lib "github.com/truauth/truauth/pkg/auth-lib"
)

// HandleGrantType grant type handler
func (req *Defaults) HandleGrantType(request *lib.AuthTokenRequest) (*lib.AccessTokenResponse, *lib.AuthErrorResponse) {
	if request.GrantType == "authorization_code" {
		return req.handleGrantAuthorizationCode(request)
	} else if request.GrantType == "refresh_token" {
		return req.handleGrantRefreshToken(request)
	}

	return nil, lib.CreateError(lib.UnauthorizedClient, "invalid access grant")
}

func (req *Defaults) handleGrantAuthorizationCode(request *lib.AuthTokenRequest) (*lib.AccessTokenResponse, *lib.AuthErrorResponse) {
	decodedCodeToken := lib.DecodeCodeJWT(req.Environment.JWTSecret, request.Code)

	authError := request.Validate(decodedCodeToken)
	if authError != nil {
		return nil, authError
	}

	// ~ access token creation
	duration := req.Configuration.ExpireDuration
	accessToken := lib.CreateAccessToken(request.ClientID, decodedCodeToken.UserID, duration, decodedCodeToken.Scope)
	signedToken := accessToken.ToJWT(request.ClientSecret)

	// refresh token creation
	refreshToken := lib.CreateRefreshToken(request.ClientID, decodedCodeToken.UserID, decodedCodeToken.Scope)
	signedRefreshToken := refreshToken.ToJWT(req.Environment.JWTSecret)

	return lib.CreateAccessTokenResponse(signedToken, lib.BearerToken, duration, signedRefreshToken), nil
}

func (req *Defaults) handleGrantRefreshToken(request *lib.AuthTokenRequest) (*lib.AccessTokenResponse, *lib.AuthErrorResponse) {
	if request.RefreshToken == "" {
		lib.CreateError(lib.InvalidRequest, "refresh token is requried for the refresh token grant")
	}

	refreshToken := lib.DecodeSignedRefreshToken(req.Environment.JWTSecret, request.RefreshToken)
	if refreshToken == nil {
		lib.CreateError(lib.InvalidToken, "invalid token")
	}

	// ~ access token creation
	duration := req.Configuration.ExpireDuration
	accessToken := lib.CreateAccessToken(request.ClientID, refreshToken.ClientID, duration, refreshToken.Scope)
	signedToken := accessToken.ToJWT(req.Environment.JWTSecret)

	return lib.CreateAccessTokenResponse(signedToken, lib.BearerToken, duration, request.RefreshToken), nil
}
