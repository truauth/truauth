package helpers

import (
	"net/http"

	grpcIdentity "github.com/truauth/truauth/pkg/grpc-identity"
)

// MarshalCredentialsFromForm creates a credentials object based of the provided form values
func MarshalCredentialsFromForm(req *http.Request) *grpcIdentity.UserIdentityRequest {
	req.ParseForm()

	return &grpcIdentity.UserIdentityRequest{
		Username: req.FormValue("username"),
		Password: req.FormValue("password"),
	}
}
