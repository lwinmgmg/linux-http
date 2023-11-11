package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"golang.org/x/exp/slices"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/lwinmgmg/linux-http/models"
	"github.com/lwinmgmg/linux-http/services"
	"github.com/lwinmgmg/linux-http/utils"
)

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

func ValidateJwtToken(tkn, key string, claim jwt.Claims) error {
	_, err := jwt.ParseWithClaims(tkn, claim, func(token *jwt.Token) (interface{}, error) {
		iss, err := token.Claims.GetIssuer()
		if err != nil {
			return nil, err
		}
		if !slices.Contains(Env.LH_ISSUERS, iss) {
			return nil, errors.New("unknown issuer")
		}
		issuedAt, err := token.Claims.GetIssuedAt()
		if err != nil {
			return nil, err
		}
		return utils.Hash256(fmt.Sprintf("%v%v", key, issuedAt.Unix())), nil
	})
	return err
}

func JwtMiddleware(tknMap map[string]int) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tknType := "Bearer"
		keyString := ctx.Request.Header.Get("Authorization")
		tknStr, err := ParseToken(keyString, tknType)
		if err != nil {
			if err == utils.ErrNotFound {
				panic(NewPanic(http.StatusUnauthorized, 1, "Authorization Required"))
			}
			if err == utils.ErrInvalid {
				panic(NewPanic(http.StatusUnauthorized, 2, "Wrong Token Type"))
			}
		}
		if Env.LH_TKN_LIMIT != 0 {
			_, ok := tknMap[tknStr]
			if ok {
				panic(NewPanic(http.StatusUnauthorized, 5, "Authorization Required!"))
			} else {
				tknMap[tknStr] = 1
				if err := services.DB.Create(models.Key{
					Key:   tknStr,
					Count: tknMap[tknStr],
				}).Error; err != nil {
					panic(err)
				}
			}
		}
		claim := jwt.RegisteredClaims{}
		if err := ValidateJwtToken(tknStr, Env.LH_SECRET, &claim); err != nil {
			if errors.Is(err, jwt.ErrTokenExpired) {
				panic(NewPanic(http.StatusUnauthorized, 3, "Authorization Required!"))
			} else {
				panic(NewPanic(http.StatusUnauthorized, 4, "Authorization Required!"))
			}
		}
	}
}
