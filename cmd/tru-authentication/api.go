package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/truauth/truauth/cmd/tru-authentication/helpers"

	lib "github.com/truauth/truauth/pkg/auth-lib"
	grpcAuthorization "github.com/truauth/truauth/pkg/grpc-authorization"
	grpcIdentity "github.com/truauth/truauth/pkg/grpc-identity"
	webserve "github.com/truauth/truauth/pkg/web-serve"
)

// AuthPage serves the auth page
func (req *Defaults) AuthPage(ctx *gin.Context) {
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

	pageTemplate := req.SitePages.RetrieveFile("login").Template
	pageTemplate.Execute(ctx.Writer, nil)
}

// CreateAuthCode creates an auth code
func (req *Defaults) CreateAuthCode(ctx *gin.Context) {
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

	// check & resolve the identity
	if resp, err := req.IdentityClient.Client.ValidateUserIdentity(ctx, credentials); err != nil || !resp.GetSuccess() {
		pageTemplate := req.SitePages.RetrieveFile("login").Template

		descript := "Error: Invalid Credentials"

		if err != nil {
			descript = "An Internal Error has Occurred While Logging In"
		}

		pageTemplate.Execute(ctx.Writer, webserve.OutTemplate{
			Error:            true,
			ErrorDescription: descript,
			Script:           webserve.DevError(err.Error()),
		})
		return
	}

	token := lib.CreateCodeToken(authRequest, credentials.GetUsername()).ToJWT(req.Environment.JWTSecret)

	grpcAuthRequest := &grpcAuthorization.AuthorizeUserRequest{
		UserID:   credentials.Username,
		ClientID: authRequest.ClientID,
	}

	// check if user has accepted agreement on conditions page
	// returns condition page if not.
	if resp, err := req.AuthorizationClient.Client.IsUserAuthorized(ctx, grpcAuthRequest); err != nil || !resp.GetSuccess() {
		temp := req.SitePages.RetrieveFile("conditions")
		if temp == nil {
			ctx.JSON(http.StatusInternalServerError, lib.CreateError("internal server error", "conditions page not found"))
			return
		}

		temp.Template.Execute(ctx.Writer, webserve.OutTemplate{
			Script: webserve.WriteSessionStorage("_truauth.clientIntro", token),
		})

		return
	}

	redirect := fmt.Sprintf("%s?code=%s&state=%s", authRequest.RedirectURI, token, authRequest.State)
	http.Redirect(ctx.Writer, ctx.Request, redirect, http.StatusFound)
}

// CreateAuthToken creates an auth token
func (req *Defaults) CreateAuthToken(ctx *gin.Context) {
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
func (req *Defaults) AuthTokenIntrospection(ctx *gin.Context) {
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

// HandleAcceptConditions handles the user accepting the conditions
func (req *Defaults) HandleAcceptConditions(ctx *gin.Context) {
	token, exists := ctx.GetQuery("t")
	authRequest, err := lib.InitAuthCode(ctx.Request)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	if !exists {
		ctx.JSON(http.StatusBadRequest, lib.CreateError("bad request", "token query is not present in request"))
		return
	}

	decodedCodeToken := lib.DecodeCodeJWT(req.Environment.JWTSecret, token)
	if decodedCodeToken.ClientID == "" || decodedCodeToken.UserID == "" {
		ctx.JSON(http.StatusBadRequest, lib.CreateError("bad request", "invalid jtw body"))
		return
	}

	updatedToken := decodedCodeToken.ToJWT(req.Environment.JWTSecret) // don't really like this.

	grpcAuthRequest := &grpcAuthorization.AuthorizeUserRequest{
		UserID:   decodedCodeToken.UserID,
		ClientID: authRequest.ClientID,
	}

	if resp, err := req.AuthorizationClient.Client.AuthorizeUser(ctx, grpcAuthRequest); err != nil || !resp.GetSuccess() {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	redirect := fmt.Sprintf("%s?code=%s&state=%s", authRequest.RedirectURI, updatedToken, authRequest.State)
	http.Redirect(ctx.Writer, ctx.Request, redirect, http.StatusFound)
}
