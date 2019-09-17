package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	lib "github.com/truauth/truauth/pkg/auth-lib"
	grpcAuthorization "github.com/truauth/truauth/pkg/grpc-authorization"
	"github.com/truauth/truauth/pkg/utilities"
	webserve "github.com/truauth/truauth/pkg/web-serve"
)

// HandleSSO extracts the sso token from the session
func (req *Defaults) HandleSSO(ctx *gin.Context, authRequest *lib.AuthCodeRequest) bool {
	cookie, err := ctx.Request.Cookie(req.SessionID)
	if err == nil {
		decodedSSO := lib.DecodeSignedSSOToken(req.Environment.JWTSecret, cookie.Value)

		grpcAuthRequest := &grpcAuthorization.AuthorizeUserRequest{
			UserID:   decodedSSO.UserID,
			ClientID: authRequest.ClientID,
		}

		if resp, err := req.AuthorizationClient.Client.IsUserAuthorized(ctx, grpcAuthRequest); err != nil || !resp.GetSuccess() {
			temp := req.SitePages.RetrieveFile("conditions")
			if temp == nil {
				ctx.JSON(http.StatusInternalServerError, lib.CreateError("internal server error", "conditions page not found"))
				return true
			}

			temp.Template.Execute(ctx.Writer, webserve.OutTemplate{
				Script: webserve.WriteSessionStorage("_truauth.clientIntro", decodedSSO.Token),
			})

			return true
		}

		redirect := fmt.Sprintf("%s?token=%s&state=%s&sso=true", authRequest.RedirectURI, decodedSSO.Token, authRequest.State)
		http.Redirect(ctx.Writer, ctx.Request, redirect, http.StatusFound)

		return true
	}

	return false
}

// CreateSSOToken creates an sso token
func (req *Defaults) CreateSSOToken(ctx *gin.Context, token string, userID string) {
	ssoToken := lib.CreateSSOToken(token, userID)

	cookie := utilities.CreateCookie(req.SessionID, "localhost:4820", "/auth", time.Now().AddDate(0, 0, 90), ssoToken.ToJWT(req.Environment.JWTSecret))

	http.SetCookie(ctx.Writer, cookie)
}
