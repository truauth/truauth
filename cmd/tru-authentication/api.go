package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/truauth/truauth/cmd/tru-authentication/helpers"

	lib "github.com/truauth/truauth/pkg/auth-lib"
	grpcIdentity "github.com/truauth/truauth/pkg/grpc-identity"
	webserve "github.com/truauth/truauth/pkg/web-serve"
)

// AuthPage serves the auth page
func (req *Request) AuthPage(ctx *gin.Context) {
	authRequest, authError := lib.InitAuthCode(ctx.Request)

	if authError != nil {
		ctx.JSON(http.StatusUnauthorized, authError)
		return
	}

	validateErr := authRequest.Validate()
	if validateErr != nil {
		ctx.JSON(http.StatusUnauthorized, validateErr)
		return
	}

	pageTemplate := req.SitePages.RetrieveFile("auth").Template
	pageTemplate.Execute(ctx.Writer, nil)
}

// CreateAuthCode creates an auth code
func (req *Request) CreateAuthCode(ctx *gin.Context) {
	authRequest, err := lib.InitAuthCode(ctx.Request)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	validateErr := authRequest.Validate()
	if validateErr != nil {
		ctx.JSON(http.StatusBadRequest, validateErr)
		return
	}

	credentials := helpers.MarshalCredentialsFromForm(ctx.Request)
	if resp, err := req.IdentityClient.Client.ValidateUserIdentity(ctx, credentials); err != nil || !resp.GetSuccess() {
		pageTemplate := req.SitePages.RetrieveFile("auth").Template

		pageTemplate.Execute(ctx.Writer, webserve.ErrorTemplate{
			Error:            true,
			ErrorDescription: "invalid credentials",
		})
		return
	}

	token := lib.CreateCodeToken(authRequest, credentials.GetUsername()).ToJWT(req.Environment.JWTSecret)

	redirect := fmt.Sprintf("%s?code=%s&state=%s", authRequest.RedirectURI, token, authRequest.State)
	http.Redirect(ctx.Writer, ctx.Request, redirect, http.StatusFound)
}

// CreateAuthToken creates an auth token
func (req *Request) CreateAuthToken(ctx *gin.Context) {
	token, err := lib.InitAuthTokenRequest(ctx.Request)
	if exit := helpers.CheckResponseError(ctx, err, token.RedirectURI, ""); exit {
		return
	}

	grantResponse, grantErr := req.HandleGrantType(token)
	if exit := helpers.CheckResponseError(ctx, grantErr, token.RedirectURI, ""); exit {
		return
	}

	ctx.JSON(http.StatusOK, grantResponse)
}

// AuthTokenIntrospection builds and returns a safe introspection view
func (req *Request) AuthTokenIntrospection(ctx *gin.Context) {
	introspectReq, err := lib.InitAccessTokenIntrospectionRequest(ctx.Request)
	if exit := helpers.CheckResponseError(ctx, err, "", ""); exit {
		return
	}

	grpcClient := &grpcIdentity.ClientIdentityRequest{
		ID: introspectReq.ClientID,
	}

	client, rpcErr := req.IdentityClient.Client.EnquireClientIdentity(ctx, grpcClient)
	if exit := helpers.CheckResponseError(ctx, &lib.AuthErrorResponse{
		Error:            lib.InvalidClient,
		ErrorDescription: rpcErr.Error(),
	}, "", ""); exit {
		return
	}

	// ensure secretes match
	validateErr := lib.CreateCustomError(client.Secret != introspectReq.ClientSecret, lib.InvalidRequest, "invalid client secret")
	if exit := helpers.CheckResponseError(ctx, validateErr, "", ""); exit {
		return
	}

	authToken := lib.DecodeSignedAccessJWT(introspectReq.ClientSecret, introspectReq.Token)
	jtwValidateErr := lib.CreateCustomError(authToken == nil, lib.InvalidClient, "client authentication was invalid")
	if exit := helpers.CheckResponseError(ctx, jtwValidateErr, "", ""); exit {
		return
	}

	ctx.JSON(http.StatusOK, authToken.MapIntrospectionResponse())
}
