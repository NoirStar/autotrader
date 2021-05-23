package restapi

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/noirstar/autotrading/backend/utils/env"
	"github.com/noirstar/autotrading/backend/utils/jwt"
	"github.com/noirstar/autotrading/backend/utils/myerr"
)

var baseURL = env.GetEnv("UPBIT_BASE_URL")

// GetAccount 전체 계좌 조회
func GetAccount(accessKey string, secretKey string) []byte {

	reqURL := baseURL + "/v1/accounts"
	tokenString := jwt.GetJwtToken(accessKey, secretKey)

	return RequestToServer(reqURL, "GET", tokenString, nil)
}

// GetOrderChance 주문 가능 정보 - 마켓별 주문 가능 정보 확인
func GetOrderChance(accessKey string, secretKey string, query interface{}) []byte {

	reqURL := baseURL + "/v1/orders/chance"
	tokenString := jwt.GetJwtTokenWithQuery(accessKey, secretKey, query)

	return RequestToServer(reqURL, "GET", tokenString, query)
}

// RequestToServer 업비트 서버로 요청
func RequestToServer(reqURL string, method string, tokenString string, query map[string]string) []byte {

	client := &http.Client{}
	req, err := http.NewRequest(method, reqURL, nil)
	myerr.CheckErr(err)

	q := url.Values{}
	for key, value := range query {
		q.Add(key, value)
	}
	req.URL.RawQuery = q.Encode()

	fmt.Println(req.URL.String())

	req.Header.Add("Authorization", "Bearer "+tokenString)
	res, err := client.Do(req)
	myerr.CheckErr(err)

	bytes, err := ioutil.ReadAll(res.Body)
	myerr.CheckErr(err)

	defer res.Body.Close()

	return bytes
}
