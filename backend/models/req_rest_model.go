package models

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/noirstar/autotrading/backend/utils/myerr"
)

func GetJwtToken(accessKey string, secretKey string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"access_key": accessKey,
		"nonce":      uuid.New(),
	})
	tokenString, err := token.SignedString([]byte(secretKey))
	myerr.CheckErr(err)
	return tokenString
}
