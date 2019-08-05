package main

import (
	"flag"
	"fmt"

	"github.com/gin-gonic/gin"

	grpcservices "github.com/truauth/truauth/cmd/tru-authentication/grpcservices"
	"github.com/truauth/truauth/pkg/middleware"
	"github.com/truauth/truauth/pkg/settings"
	webserve "github.com/truauth/truauth/pkg/web-serve"
)

// Defaults struct used to pass request information
type Defaults struct {
	SitePages      webserve.Files
	IdentityClient *grpcservices.IdentityClient
	Environment    *settings.Environment
	Configuration  *settings.Configuration
}

func main() {
	httpPort := flag.String("p", "4820", "specificed port to start the http server on")

	config := &settings.Configuration{}
	settings.Init(config, "configuration")

	env := &settings.Environment{}
	settings.Init(env, "env")

	gRPCIdentityClient := grpcservices.CreateIdentityGRPCClient(config.IdentityServiceAddress)

	req := &Defaults{
		SitePages:      webserve.Init("auth"),
		IdentityClient: gRPCIdentityClient,
		Environment:    env,
		Configuration:  config,
	}

	router := gin.Default()

	router.GET("/auth", req.AuthPage)
	router.POST("/code", req.CreateAuthCode)
	router.POST("/token", middleware.Register(req.CreateAuthToken, middleware.EnableCORS))
	router.POST("/token-info", middleware.Register(req.AuthTokenIntrospection, middleware.EnableCORS))

	fmt.Println("Server Started on Port 4820")
	router.Run(*httpPort)
}
