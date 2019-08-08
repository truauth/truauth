package grpcservices

import (
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"

	grpcIdentity "github.com/truauth/truauth/pkg/grpc-identity"
)

// IdentityClient gRPC client for the identity service
type IdentityClient struct {
	Client grpcIdentity.IdentityClient
}

// CreateIdentityGRPCClient creates a grpc identity client
func CreateIdentityGRPCClient(hostAddress string) *IdentityClient {
	k := keepalive.ClientParameters{
		Time:                time.Minute * 2,
		PermitWithoutStream: true,
	}

	cc, err := grpc.Dial(hostAddress, grpc.WithInsecure(), grpc.WithKeepaliveParams(k))
	if err != nil {
		panic(err)
	}

	return &IdentityClient{
		Client: grpcIdentity.NewIdentityClient(cc),
	}
}
