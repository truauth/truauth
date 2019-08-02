package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	lib "github.com/truauth/truauth/pkg/auth-lib"
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

}

// CreateAuthToken creates an auth token
func (req *Request) CreateAuthToken(ctx *gin.Context) {

}

// AuthTokenIntrospection builds and returns a safe introspection view
func (req *Request) AuthTokenIntrospection(ctx *gin.Context) {

}
