package token

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
)

var InvalidTokenError = errors.New("invalid token")

func NewToken(issuer, audience, subject string, expiration time.Duration, signKey []byte) (string, error) {
	claims :=
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(expiration).Unix(),
			Issuer:    issuer,
			Audience:  audience,
			Subject:   subject,
		}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(signKey)
}

func ParseToken(issuer, tokenValue string, signKey []byte) (*jwt.StandardClaims, error) {
	token, err := jwt.ParseWithClaims(
		tokenValue,
		&jwt.StandardClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return signKey, nil
		})
	if err != nil {
		return nil, errors.Wrap(err, "invalid token")
	}
	claims, ok := token.Claims.(*jwt.StandardClaims)
	if !ok || !token.Valid {
		return nil, errors.WithStack(InvalidTokenError)
	}
	if claims.Issuer != issuer {
		return nil, errors.WithStack(InvalidTokenError)
	}
	return claims, nil
}
