package main

import (
	"context"

	"golang.org/x/crypto/bcrypt"

	"github.com/truauth/truauth/cmd/tru-identity/db"
	grpcIdentity "github.com/truauth/truauth/pkg/grpc-identity"
	mapping "github.com/truauth/truauth/pkg/identity-mapping"
)

// EnquireUserIdentity enquires about a user, retuns the found user (in a safe response) if found.
func (server *ServiceRequest) EnquireUserIdentity(ctx context.Context, request *grpcIdentity.UserIdentityRequest) (*grpcIdentity.SafeUserIdentity, error) {
	pgUserReq := &db.UserDbRequest{server.PGCreds}
	user, err := pgUserReq.FindByUsername(request.Username)
	if err != nil {
		return nil, err
	}

	safeUser := mapping.FromUnsafeUserToSafe(user)

	return safeUser, nil
}

// RegisterUserIdentity registers a user identity
func (server *ServiceRequest) RegisterUserIdentity(ctx context.Context, request *grpcIdentity.UnsafeUserIdentity) (*grpcIdentity.SuccessResponse, error) {
	pgUserReq := &db.UserDbRequest{server.PGCreds}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	request.Password = string(hashedPassword) // possibly rething this mutation?

	err = pgUserReq.DirectCreate(request)

	return &grpcIdentity.SuccessResponse{
		Success: err == nil,
	}, err
}

// ValidateUserIdentity validates a user identity, returns success response if passed
func (server *ServiceRequest) ValidateUserIdentity(ctx context.Context, request *grpcIdentity.UserIdentityRequest) (*grpcIdentity.SuccessResponse, error) {
	pgUserReq := &db.UserDbRequest{server.PGCreds}
	user, err := pgUserReq.FindByUsername(request.Username)
	if err != nil {
		return nil, err
	}

	if request.GetPassword() == user.GetPassword() { // some parts of the dev system are not using bcrypt. (todo: disable for prod)
		return &grpcIdentity.SuccessResponse{
			Success: true,
		}, nil
	}

	bcryptErr := bcrypt.CompareHashAndPassword([]byte(user.GetPassword()), []byte(request.GetPassword()))

	return &grpcIdentity.SuccessResponse{
		Success: bcryptErr == nil,
	}, bcryptErr
}
