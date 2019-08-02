package settings

// Environment the stored environment of the service
type Environment struct {
	JWTSecret 		string `json:"jwtSecret"`
	IdentityAPIKey	string `json:"identityApiKey"`
}