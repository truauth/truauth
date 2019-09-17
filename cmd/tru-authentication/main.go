package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"


	grpcservices "github.com/truauth/truauth/cmd/tru-authentication/grpcservices"
	"github.com/truauth/truauth/pkg/middleware"
	"github.com/truauth/truauth/pkg/settings"
	"github.com/truauth/truauth/pkg/utilities"
	webserve "github.com/truauth/truauth/pkg/web-serve"
)

// Defaults struct used to pass request information
type Defaults struct {
	SitePages           webserve.Files
	IdentityClient      *grpcservices.IdentityClient
	AuthorizationClient *grpcservices.AuthorizationClient
	Environment         *settings.Environment
	Configuration       *settings.Configuration
	SessionID           string
}
func init() {
	_ = godotenv.Load("configuration.env", "default.env", "postgres.env")
}
func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "4820"
	}

	httpPort := flag.String("p", ":"+port, "specificed port to start the http server on")
	req := initDefaults()

	router := gin.Default()

	router.GET("/auth", req.AuthPage)
	router.POST("/code", req.CreateAuthCode)
	router.POST("/token", middleware.Register(req.CreateAuthToken, middleware.EnableCORS))
	router.POST("/token-info", middleware.Register(req.AuthTokenIntrospection, middleware.EnableCORS))
	router.POST("/approve-conditions", middleware.Register(req.HandleAcceptConditions, middleware.EnableCORS))

	fmt.Println(fmt.Sprintf("Server started on port %s", *httpPort))
	router.Run(*httpPort)
}

func initDefaults() *Defaults {
	tokenLifetime := os.Getenv("TOKEN_LIFETIME")
	tokenLifetimeInt, err := strconv.Atoi(tokenLifetime)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	config := &settings.Configuration{
		ExpireDuration:              tokenLifetimeInt,
		IdentityServiceAddress:      os.Getenv("IDENTITY_SERVICE_ADDRESS"),
		AuthorizationServiceAddress: os.Getenv("AUTHORIZATION_SERVICE_ADDRESS"),
	}
	env := &settings.Environment{
		JWTSecret:      os.Getenv("JWT_SECRET"),
		IdentityAPIKey: os.Getenv("IDENTITY_API_KEY"),
	}

	gRPCIdentityClient := grpcservices.CreateIdentityGRPCClient(config.IdentityServiceAddress)
	gRPCAuthorizationClient := grpcservices.CreateAuthorizationGRPCClient(config.AuthorizationServiceAddress)

	return &Defaults{
		SitePages:           webserve.Init("login", "conditions"),
		IdentityClient:      gRPCIdentityClient,
		AuthorizationClient: gRPCAuthorizationClient,
		Environment:         env,
		Configuration:       config,
		SessionID:           utilities.UUID(),
	}
}
