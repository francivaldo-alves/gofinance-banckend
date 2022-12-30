package util

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"strings"
)

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func ValidateToken(ctx *gin.Context, token string) error {
	claims := &Claims{}
	var jwtSignedKey = []byte("secret_key")
	tokenParse, err := jwt.ParseWithClaims(token, claims,
		func(t *jwt.Token) (interface{}, error) {
			return jwtSignedKey, nil
		})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			ctx.JSON(http.StatusUnauthorized, err)
			return err
		}
		ctx.JSON(http.StatusBadRequest, err)
		return err
	}
	if !tokenParse.Valid {
		ctx.JSON(http.StatusUnauthorized, "Token invalido")
		return nil
	}
	ctx.Next()
	return nil
}

func GetTokenInHeaderAndVerify(ctx *gin.Context) error {
	authorizationHanderKey := ctx.GetHeader("authorization")
	filds := strings.Fields(authorizationHanderKey)
	tokenTovalidate := filds[1]
	errOnValidateToken := ValidateToken(ctx, tokenTovalidate)
	if errOnValidateToken != nil {
		return errOnValidateToken

	}
	return nil
}
