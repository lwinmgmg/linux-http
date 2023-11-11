package utils

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GetToken(issuer, tokenKey string, expiresAfter time.Duration) string {
	nowTime := time.Now().UTC()
	claim := jwt.RegisteredClaims{
		Issuer:    issuer,
		IssuedAt:  jwt.NewNumericDate(nowTime),
		ExpiresAt: jwt.NewNumericDate(nowTime.Add(expiresAfter)),
	}
	tkn := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	output, _ := tkn.SignedString(Hash256(fmt.Sprintf("%v%v", tokenKey, claim.IssuedAt.Unix())))
	return output
}
