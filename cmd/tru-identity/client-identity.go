package main

import (
	"context"

	"github.com/truauth/truauth/cmd/tru-identity/db"
	"github.com/truauth/truauth/pkg/grpc-identity"
)

// EnquireClientIdentity enquires information about the client identity
func (server *ServiceServer) EnquireClientIdentity(ctx context.Context, request *grpcIdentity.ClientIdentityRequest) (*grpcIdentity.ClientIdentity, error) {
	pgClientReq := &db.ClientDbRequest{PGCreds: server.PGCreds}
	client, err := pgClientReq.FindByID(request.ID)

	if err != nil {
		return nil, err
	}

	return client, nil
}

// RegisterClientIdentity registers a client provided the client identity
func (server *ServiceServer) RegisterClientIdentity(ctx context.Context, request *grpcIdentity.ClientIdentity) (*grpcIdentity.CreatedResponse, error) {
	pgClientReq := &db.ClientDbRequest{PGCreds: server.PGCreds}
	err := pgClientReq.DirectCreate(request)

	return &grpcIdentity.CreatedResponse{
		Success: err != nil,
	}, err
}