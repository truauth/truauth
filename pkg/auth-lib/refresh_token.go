package lib

import (
	"github.com/dgrijalva/jwt-go"
	"time"
	"fmt"
)

// RefreshToken auth refresh token assigned with the bearer token
type RefreshToken struct {
	IAT 	int64
	UserID 	string
	ClientID string
	Scope 	string
}

// CreateRefreshToken creates a new refresh token
func CreateRefreshToken(clientID string, userID string, scope string) *RefreshToken {
	return &RefreshToken{
		IAT: time.Now().Unix(),
		UserID: userID,
		ClientID: clientID,
		Scope: scope,
	}
}

// ToJWT creates a jwt token from the code token
func (request *RefreshToken) ToJWT(clientSecret string) string {
	token := jwt.New(jwt.SigningMethodHS256)
	
	token.Claims = jwt.MapClaims{
		"client_id": request.ClientID,
		"user_id": request.UserID,
		"iat": request.IAT,
		"scope": request.Scope,
	}
	
	tokenString, _ := token.SignedString([]byte(clientSecret))
	
	return tokenString
}

// DecodeSignedRefreshToken decodes the provided jtw token
func DecodeSignedRefreshToken(secret string, signedToken string) *CodeToken {
	claims := jwt.MapClaims{} 
	token, _ := jwt.ParseWithClaims(signedToken, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if !token.Valid {
		return nil
	}

	iat := claims["iat"].(float64)

	return &CodeToken{
		IAT: int64(iat),
		ClientID: fmt.Sprint(claims["clientID"]),
		UserID: fmt.Sprint(claims["userID"]),
		Scope: fmt.Sprint(claims["scope"]),
	}
}
