package lib

import (
	"github.com/dgrijalva/jwt-go"
	"time"
	"fmt"
)

// AccessToken access token struct used to decode / encode jtw access token
type AccessToken struct {
	ClientID		string
	ExpiresAt 		int64
	UserID			string
	IAT				int64
	Scope 	string
}

// CreateAccessToken creates an access token
func CreateAccessToken(clientID string, userID string, expiresIn int, scope string) *AccessToken {
	return &AccessToken{
		ClientID: clientID,
		UserID: userID,
		IAT: time.Now().Unix(),
		ExpiresAt: time.Now().Add(time.Second * time.Duration(expiresIn)).Unix(),
		Scope: scope,
	}
}

// ToJWT creates a jtw token from the access token struct
func (request *AccessToken) ToJWT(clientSecret string) string {
	token := jwt.New(jwt.SigningMethodHS256)

	token.Claims = jwt.MapClaims{
		"client_id": request.ClientID,
		"expires_at": request.ExpiresAt,
		"user_id": request.UserID,
		"iat": request.IAT,
		"scope": request.Scope,
	}

	tokenString, _ := token.SignedString([]byte(clientSecret))

	return tokenString
}

// DecodeSignedAccessJWT decodes the provided jtw token
func DecodeSignedAccessJWT(secret string, signedToken string) *AccessToken {
	claims := jwt.MapClaims{} 
	token, _ := jwt.ParseWithClaims(signedToken, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if !token.Valid {
		return nil
	}

	expiresAt := claims["expires_at"].(float64)
	iat := claims["iat"].(float64)

	return &AccessToken{
		ClientID: fmt.Sprint(claims["client_id"]),
		UserID: fmt.Sprint(claims["user_id"]),
		ExpiresAt: int64(expiresAt),
		IAT: int64(iat),
		Scope: fmt.Sprint(claims["scope"]),
	}
}

// HasExpired determines if the code token is expired
func (request *AccessToken) HasExpired() bool {
	currentTime := time.Now().Unix()
	return currentTime > request.ExpiresAt
}