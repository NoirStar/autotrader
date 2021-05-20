package models

import (
	"crypto/sha512"
	"fmt"
	"net/url"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/noirstar/autotrading/backend/utils/myerr"
)

// GetJwtToken jwt token 생성
func GetJwtToken(accessKey string, secretKey string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"access_key": accessKey,
		"nonce":      uuid.New(),
	})
	tokenString, err := token.SignedString([]byte(secretKey))
	myerr.CheckErr(err)
	return tokenString
}

// GetJwtTokenWithQuery jwt token 생성 (쿼리값 포함)
func GetJwtTokenWithQuery(accessKey string, secretKey string, query map[string]string) string {

	params := url.Values{}
	for key, value := range query {
		params.Add(key, value)
	}

	queryHash := fmt.Sprintf("%x", sha512.Sum512([]byte(params.Encode())))

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"access_key":     accessKey,
		"nonce":          uuid.New(),
		"query_hash":     queryHash,
		"query_hash_alg": "SHA512",
	})
	tokenString, err := token.SignedString([]byte(secretKey))
	myerr.CheckErr(err)
	return tokenString
}
