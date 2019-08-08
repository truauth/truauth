package grpcservices

import (
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"

	grpcAuthorization "github.com/truauth/truauth/pkg/grpc-authorization"
)

// AuthorizationClient gRPC client for the authorization service
type AuthorizationClient struct {
	Client grpcAuthorization.AuthorizationClient
}

// CreateAuthorizationGRPCClient creates a grpc authorization client
func CreateAuthorizationGRPCClient(hostAddress string) *AuthorizationClient {
	k := keepalive.ClientParameters{
		Time:                time.Minute * 2,
		PermitWithoutStream: true,
	}

	cc, err := grpc.Dial(hostAddress, grpc.WithInsecure(), grpc.WithKeepaliveParams(k))
	if err != nil {
		panic(err)
	}

	return &AuthorizationClient{
		Client: grpcAuthorization.NewAuthorizationClient(cc),
	}
}
