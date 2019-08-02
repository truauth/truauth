package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/truauth/truauth/pkg/lib"
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
	ctx.

}