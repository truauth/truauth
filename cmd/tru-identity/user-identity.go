package main

import (
	"context"
	
	"github.com/truauth/truauth/pkg/grpc-identity"
	"github.com/truauth/truauth/pkg/identity-mapping"
	"github.com/truauth/truauth/cmd/tru-identity/db"
)

// EnquireUserIdentity enquires about a user, retuns the found user (in a safe response) if found.
func (server *ServiceServer) EnquireUserIdentity(ctx context.Context, request *grpcIdentity.UserIdentityRequest) (*grpcIdentity.SafeUserIdentity, error) {
	pgUserReq := &db.UserDbRequest{server.PGCreds}
	user, err := pgUserReq.FindByUsername(request.Username)
	if err != nil {
		return nil, err
	}

	safeUser := mapping.FromUnsafeUserToSafe(user)

	return safeUser, nil
} 

// RegisterUserIdentity registers a user identity
func (server *ServiceServer) RegisterUserIdentity(ctx context.Context, request *grpcIdentity.UnsafeUserIdentity) (*grpcIdentity.CreatedResponse, error) {
	pgUserReq := &db.UserDbRequest{server.PGCreds}
	err := pgUserReq.DirectCreate(request)

	return &grpcIdentity.CreatedResponse{
		Success: err != nil,
	}, err
}
