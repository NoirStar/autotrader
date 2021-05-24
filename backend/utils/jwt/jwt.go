package jwt

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
func GetJwtTokenWithQuery(accessKey string, secretKey string, query map[string]interface{}) string {

	q := url.Values{}

	for key, value := range query {
		switch val := value.(type) {
		case string:
			q.Add(key, value.(string))
		case int, uint32, uint64:
			q.Add(key, value.(string))
		case []string:
			for _, v := range val {
				q.Add(key, v)
			}
		case []interface{}:
			for _, v := range val {
				q.Add(key, v.(string))
			}
		}
	}

	queryHash := fmt.Sprintf("%x", sha512.Sum512([]byte(q.Encode())))

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
