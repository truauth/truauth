package grpcservices

import (
	"google.golang.org/grpc"

	grpcIdentity "github.com/truauth/truauth/pkg/grpc-identity"
)

// IdentityClient gRPC client for the identity service
type IdentityClient struct {
	Client grpcIdentity.IdentityClient
}

// CreateIdentityGRPCClient creates a grpc identity client
func CreateIdentityGRPCClient(hostAddress string) *IdentityClient {
	cc, err := grpc.Dial(hostAddress, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	return &IdentityClient{
		Client: grpcIdentity.NewIdentityClient(cc),
	}
}
