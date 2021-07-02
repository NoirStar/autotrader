package utils

import (
	"crypto/sha512"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// CreateUpbitJwt jwt token 생성
func CreateUpbitJwt(accessKey string, secretKey string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"access_key": accessKey,
		"nonce":      uuid.New(),
	})
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// CreateUpbitJwtQuery jwt token 생성 (쿼리값 포함)
func CreateUpbitJwtQuery(accessKey string, secretKey string, query map[string]interface{}) (string, error) {

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
	if err != nil {
		return "", nil
	}
	return tokenString, nil
}

// CreateJwt creates jwt token
func CreateJwt(c echo.Context, data map[string]interface{}, expire time.Time, cookieName string) (string, error) {

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	for key, val := range data {
		claims[key] = val
	}
	claims["exp"] = expire.Unix()

	tokenString, err := token.SignedString([]byte(GetEnv("JWT_SECRET_KEY")))
	if err != nil {
		return "", err
	}

	if cookieName != "" {
		cookie := new(http.Cookie)
		cookie.Name = cookieName
		cookie.Value = tokenString
		cookie.Expires = expire
		c.SetCookie(cookie)
	}
	return tokenString, nil
}

// GenerateJwt makes access token
func GenerateJwt(c echo.Context, claims map[string]interface{}) (string, error) {
	// access 토큰 생성: 유효기간 20분
	accessToken, err := CreateJwt(
		c,
		claims,
		time.Now().Add(time.Minute*20),
		"access_token",
	)

	if err != nil {
		return "", err
	}

	return accessToken, nil
}
