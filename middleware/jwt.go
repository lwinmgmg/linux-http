package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/lwinmgmg/linux-http/utils"
)

func GetToken(issuer, tokenKey string, expiresAfter time.Duration) string {
	nowTime := time.Now().UTC()
	claim := jwt.RegisteredClaims{
		Issuer:    issuer,
		IssuedAt:  jwt.NewNumericDate(nowTime),
		ExpiresAt: jwt.NewNumericDate(nowTime.Add(expiresAfter)),
	}
	tkn := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	output, _ := tkn.SignedString(utils.Hash256(fmt.Sprintf("%v%v", tokenKey, claim.IssuedAt)))
	return output
}

func ParseToken(keyString, tokenType string) (string, error) {
	if keyString == "" {
		return "", utils.ErrNotFound
	}
	inputTokenType := keyString[0:len(tokenType)]
	inputTokenString := keyString[len(tokenType):]
	if inputTokenType != tokenType {
		return "", utils.ErrInvalid
	}
	return strings.TrimSpace(inputTokenString), nil
}

func ValidateJwtToken(tkn, key, claim *jwt.Claims)

func JwtMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenType := "Bearer"
		keyString := ctx.Request.Header.Get("Authorization")
		inputTokenString, err := ParseToken(keyString, tokenType)
		if err != nil {
			if err == utils.ErrNotFound {
				panic(NewPanic(http.StatusUnauthorized, 1, "Authorization Required"))
			}
			if err == utils.ErrInvalid {
				panic(NewPanic(http.StatusUnauthorized, 2, "Wrong Token Type"))
			}
		}
		claim := jwt.RegisteredClaims{}
		if err := ValidateToken(inputTokenString, tokenKey, &claim); err != nil {
			if errors.Is(err, jwt.ErrTokenExpired) {
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, datamodels.DefaultResponse{
					Code:    3,
					Message: "Authorization Required! [TokenExpired]",
				})
			} else {
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, datamodels.DefaultResponse{
					Code:    4,
					Message: fmt.Sprintf("Authorization Required! [%v]", err),
				})
			}
			return
		}
		ctx.Set("userCode", claim.Subject)
	}
}
