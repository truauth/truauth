package lib

// AuthErrorResponse openid complinent error response
type AuthErrorResponse struct {
	Error				AuthError		`json:"error"`
	ErrorDescription 	string 			`json:"error_description"`
	State				string 			`json:"state"`
}

// CreateError creates an instance of the auth error response
func CreateError(authError AuthError, description string) *AuthErrorResponse{
	return &AuthErrorResponse{
		Error: authError,
		ErrorDescription: description,
	}
}

// CreateCustomError creates a custom error based on the paramaters
func CreateCustomError(condition bool, authError AuthError, description string) *AuthErrorResponse {
	if condition {
		return &AuthErrorResponse{
			Error: authError,
			ErrorDescription: description,
		}
	}

	return nil
}