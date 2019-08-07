package settings

// Configuration configuration struct used to store safe application settings
type Configuration struct {
	ExpireDuration              int
	IdentityServiceAddress      string
	AuthorizationServiceAddress string
}
