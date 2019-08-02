package main

import (
	"github.com/gin-gonic/gin"

	"github.com/truauth/truauth/pkg/settings"

	identitygrpc "github.com/truauth/truauth/cmd/tru-authentication/identity-grpc"
	webserve "github.com/truauth/truauth/pkg/web-serve"
)

// Request struct used to pass request information
type Request struct {
	SitePages      webserve.Files
	IdentityClient *identitygrpc.IdentityClient
	Environment    *settings.Environment
	Configuration  *settings.Configuration
}

func main() {
	gRPCIdentityClient := identitygrpc.CreateGRPCClient("locahost:4820") // todo, arg this

	req := &Request{
		SitePages:      webserve.Init("auth"),
		IdentityClient: gRPCIdentityClient,
	}

	gServe := gin.Default()

	gServe.GET("/auth", req.AuthPage)
	gServe.POST("/code", req.CreateAuthCode)
	gServe.POST("/token", req.CreateAuthToken)
	gServe.POST("/token-info", req.AuthTokenIntrospection)
}
