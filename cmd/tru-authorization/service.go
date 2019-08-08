package main

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/truauth/truauth/cmd/tru-authorization/db"
	grpcAuthorization "github.com/truauth/truauth/pkg/grpc-authorization"
)

// AuthorizeUser authorizes a single user with the desired client
func (service *ServiceRequest) AuthorizeUser(ctx context.Context, request *grpcAuthorization.AuthorizeUserRequest) (*grpcAuthorization.SuccessResponse, error) {
	dbRequest := &db.AuthDbRequest{
		PGCreds: service.PGCreds,
	}

	user, err := dbRequest.FindByUserID(request.UserID)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	if user == nil { // just create
		err = dbRequest.Create(&db.AuthorizationColumn{
			UserId:            request.UserID,
			AuthorizedClients: request.ClientID,
		})
		if err != nil {
			return nil, err
		}
	}

	if strings.Contains(user.AuthorizedClients, request.ClientID) {
		return &grpcAuthorization.SuccessResponse{
			Success: true,
		}, nil
	}

	user.AuthorizedClients = fmt.Sprintf("%s, %s", user.AuthorizedClients, request.ClientID)
	err = dbRequest.Update(user)

	return nil, err
}

// IsUserAuthorized checks if the requested user & client are authorized
func (service *ServiceRequest) IsUserAuthorized(ctx context.Context, request *grpcAuthorization.AuthorizeUserRequest) (*grpcAuthorization.SuccessResponse, error) {
	dbRequest := &db.AuthDbRequest{
		PGCreds: service.PGCreds,
	}

	user, err := dbRequest.FindByUserID(request.UserID)
	if err != nil {
		return &grpcAuthorization.SuccessResponse{
			Success: false,
		}, err
	}

	if strings.Contains(user.AuthorizedClients, request.ClientID) { // we already have it
		return &grpcAuthorization.SuccessResponse{
			Success: true,
		}, nil
	}

	return &grpcAuthorization.SuccessResponse{
		Success: false,
	}, nil
}
