package grpcservices

import (
	"google.golang.org/grpc"

	grpcAuthorization "github.com/truauth/truauth/pkg/grpc-authorization"
)

// AuthorizationClient gRPC client for the authorization service
type AuthorizationClient struct {
	Client grpcAuthorization.AuthorizationClient
}

// CreateAuthorizationGRPCClient creates a grpc authorization client
func CreateAuthorizationGRPCClient(hostAddress string) *AuthorizationClient {
	cc, err := grpc.Dial(hostAddress, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	return &AuthorizationClient{
		Client: grpcAuthorization.NewAuthorizationClient(cc),
	}
}
