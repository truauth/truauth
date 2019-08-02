package lib

// TokenType Access token type
type TokenType string

const (
	BearerToken TokenType = "bearer"
	BasicToken TokenType = "basic"
)