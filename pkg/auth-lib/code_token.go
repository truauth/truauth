package lib

import (
	"fmt"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// CodeToken auth code token request model
type CodeToken struct {
	Expire      int64 //refactor to expires_at
	UserID      string
	RedirectURI string
	Scope       string
	IAT         int64
	ClientID    string
}

// CreateCodeToken creates a code token
func CreateCodeToken(request *AuthCodeRequest, userID string) *CodeToken {
	return &CodeToken{
		Expire:      time.Now().Add(time.Second * 30).Unix(),
		UserID:      userID,
		RedirectURI: request.RedirectURI,
		Scope:       strings.Join(request.Scope, " "),
		IAT:         time.Now().Unix(),
		ClientID:    request.ClientID,
	}
}

// ToJWT creates a jwt token from the code token
func (request *CodeToken) ToJWT(secret string) string {
	token := jwt.New(jwt.SigningMethodHS256)

	token.Claims = jwt.MapClaims{
		"client_id":    request.ClientID,
		"expire":       request.Expire,
		"user_id":      request.UserID,
		"redirect_uri": request.RedirectURI,
		"scope":        request.Scope,
		"iat":          request.IAT,
	}

	tokenString, _ := token.SignedString([]byte(secret))

	return tokenString
}

// DecodeCodeJWT decodes the provided jtw token
func DecodeCodeJWT(secret string, signedToken string) *CodeToken {
	claims := jwt.MapClaims{} //switch to using a customClaim struct and nesting standard claims https://godoc.org/github.com/dgrijalva/jwt-go#StandardClaims
	token, _ := jwt.ParseWithClaims(signedToken, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if !token.Valid {
		return nil
	}

	expiresAt := claims["expire"].(float64)
	iat := claims["iat"].(float64)

	return &CodeToken{
		Expire:      int64(expiresAt),
		ClientID:    fmt.Sprint(claims["client_id"]),
		UserID:      fmt.Sprint(claims["user_id"]),
		RedirectURI: fmt.Sprint(claims["redirect_uri"]),
		Scope:       fmt.Sprint(claims["scope"]),
		IAT:         int64(iat),
	}
}

// HasExpired determines if the code token is expired
func (request *CodeToken) HasExpired() bool {
	currentTime := time.Now().Unix()
	return currentTime > request.Expire
}
