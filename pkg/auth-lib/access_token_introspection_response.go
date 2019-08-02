package lib

// AccessTokenIntrospectionResponse introspection response 
type AccessTokenIntrospectionResponse struct {
	Active bool `json:"active"`
	Scope string `json:"scope"`
	ClientID string `json:"client_id"`
	UserID string `json:"user_id"`
	Exp int64 `json:"exp"`
}

// MapIntrospectionResponse mapes the access token to introspection response
func (token *AccessToken) MapIntrospectionResponse() *AccessTokenIntrospectionResponse{
	return &AccessTokenIntrospectionResponse{
		Active: token.HasExpired(),
		Scope: token.Scope,
		ClientID: token.ClientID,
		UserID: token.UserID,
		Exp: token.ExpiresAt,
	}
}