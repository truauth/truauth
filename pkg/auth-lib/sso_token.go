package lib

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// SSOToken auth refresh token assigned with the bearer token
type SSOToken struct {
	UserID string
	IAT    int64
	Token  string
}

// CreateSSOToken creates a new refresh token
func CreateSSOToken(token string, userID string) *SSOToken {
	return &SSOToken{
		IAT:    time.Now().Unix(),
		UserID: userID,
		Token:  token,
	}
}

// ToJWT creates a jwt token from the code token
func (request *SSOToken) ToJWT(clientSecret string) string {
	token := jwt.New(jwt.SigningMethodHS256)

	token.Claims = jwt.MapClaims{
		"user_id": request.UserID,
		"iat":     request.IAT,
		"token":   request.Token,
	}

	tokenString, _ := token.SignedString([]byte(clientSecret))

	return tokenString
}

// DecodeSignedSSOToken decodes the provided jtw token
func DecodeSignedSSOToken(secret string, signedToken string) *SSOToken {
	claims := jwt.MapClaims{}
	token, _ := jwt.ParseWithClaims(signedToken, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if !token.Valid {
		return nil
	}

	iat := claims["iat"].(float64)

	return &SSOToken{
		IAT:    int64(iat),
		UserID: fmt.Sprint(claims["user_id"]),
		Token:  fmt.Sprint(claims["token"]),
	}
}
