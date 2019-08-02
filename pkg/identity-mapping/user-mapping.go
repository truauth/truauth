package mapping

import (
	"github.com/truauth/truauth/pkg/grpc-identity"
)

// FromUnsafeUserToSafe maps an unsafe user a to safe model
func FromUnsafeUserToSafe(unsafe *grpcIdentity.UnsafeUserIdentity) *grpcIdentity.SafeUserIdentity {
	return &grpcIdentity.SafeUserIdentity{
		Username: unsafe.Username,
		Email: unsafe.Email,
		Type: unsafe.Type,
	}
}