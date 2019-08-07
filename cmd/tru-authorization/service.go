package main

import (
	"context"

	"github.com/truauth/truauth/cmd/tru-authorization/db"
	grpcAuthorization "github.com/truauth/truauth/pkg/grpc-authorization"
)

// AuthorizeUser authorizes a single user with the desired client
func (service *ServiceRequest) AuthorizeUser(ctx context.Context, request *grpcAuthorization.AuthorizeUserRequest) (*grpcAuthorization.SuccessResponse, error) {
	dbRequest := &db.AuthDbRequest{
		PGCreds: service.PGCreds,
	}

	user, err := dbRequest.FindByUserID(request.UserID)
	if err != nil {
		return nil, err
	}

	if user == nil { // just create
		err = dbRequest.Create(&db.AuthorizationColumn{
			ID:                request.UserID,
			AuthorizedClients: []string{request.ClientID},
		})
		if err != nil {
			return nil, err
		}
	}

	updatedClients := make([]string, len(user.AuthorizedClients)+1)
	for idx := 0; idx < len(user.AuthorizedClients); idx++ {
		if user.AuthorizedClients[idx] == request.ClientID {
			return &grpcAuthorization.SuccessResponse{
				Success: true,
			}, nil
		}
		updatedClients[idx] = user.AuthorizedClients[idx]
	}

	updatedClients[len(user.AuthorizedClients)+1] = request.ClientID
	user.AuthorizedClients = updatedClients

	err = dbRequest.Update(user)

	return nil, err
}

// IsUserAuthorized checks if the requested user & client are authorized
func (service *ServiceRequest) IsUserAuthorized(ctx context.Context, request *grpcAuthorization.AuthorizeUserRequest) (*grpcAuthorization.SuccessResponse, error) {
	return nil, nil
}
