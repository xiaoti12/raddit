package jwt

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

const TokenExpireDuration = time.Hour * 4

var ErrorInvalidToken = errors.New("invalid token")
var (
	mySecret = []byte("TokimekiRunners")
	myIssuer = "raddit"
)

type RadditJWTClaims struct {
	UserID   int64  `json:"user_id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func GenToken(userID int64, username string) (string, error) {
	claims := &RadditJWTClaims{
		UserID:   userID,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(TokenExpireDuration)),
			Issuer:    myIssuer,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(mySecret)
}

func ParseToken(tokenStr string) (*RadditJWTClaims, error) {
	var claims = new(RadditJWTClaims)
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return mySecret, nil
	})
	if err != nil {
		return nil, err
	}
	if token.Valid {
		return claims, nil
	} else {
		return nil, ErrorInvalidToken
	}
}
