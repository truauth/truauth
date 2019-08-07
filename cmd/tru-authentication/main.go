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
	SitePages           webserve.Files
	IdentityClient      *grpcservices.IdentityClient
	AuthorizationClient *grpcservices.AuthorizationClient
	Environment         *settings.Environment
	Configuration       *settings.Configuration
}

func main() {
	httpPort := flag.String("p", "4820", "specificed port to start the http server on")
	req := initDefaults()

	router := gin.Default()

	router.GET("/auth", req.AuthPage)
	router.POST("/code", req.CreateAuthCode)
	router.POST("/token", middleware.Register(req.CreateAuthToken, middleware.EnableCORS))
	router.POST("/token-info", middleware.Register(req.AuthTokenIntrospection, middleware.EnableCORS))

	fmt.Println(fmt.Sprintf(" Server started on port %s", *httpPort))
	router.Run(*httpPort)
}

func initDefaults() *Defaults {
	config := &settings.Configuration{}
	settings.Init(config, "configuration")

	env := &settings.Environment{}
	settings.Init(env, "env")

	gRPCIdentityClient := grpcservices.CreateIdentityGRPCClient(config.IdentityServiceAddress)
	gRPCAuthorizationClient := grpcservices.CreateAuthorizationGRPCClient(config.AuthorizationServiceAddress)

	return &Defaults{
		SitePages:           webserve.Init("login"),
		IdentityClient:      gRPCIdentityClient,
		AuthorizationClient: gRPCAuthorizationClient,
		Environment:         env,
		Configuration:       config,
	}
}
