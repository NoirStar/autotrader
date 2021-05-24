package restapi

import (
	"encoding/json"
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
func GetOrderChance(accessKey string, secretKey string, query map[string]interface{}) []byte {

	reqURL := baseURL + "/v1/orders/chance"
	tokenString := jwt.GetJwtTokenWithQuery(accessKey, secretKey, query)

	return RequestToServer(reqURL, "GET", tokenString, query)
}

// GetOrderSearch 개별 주문 조회 - 주문 UUID 를 통해 개별 주문건을 조회
func GetOrderSearch(accessKey string, secretKey string, query map[string]interface{}) []byte {

	reqURL := baseURL + "/v1/order"
	tokenString := jwt.GetJwtTokenWithQuery(accessKey, secretKey, query)

	return RequestToServer(reqURL, "GET", tokenString, query)
}

// GetOrdersSearch 주문 리스트 조회 - 주문 리스트를 조회
func GetOrdersSearch(accessKey string, secretKey string, query map[string]interface{}) []byte {

	reqURL := baseURL + "/v1/orders"
	tokenString := jwt.GetJwtTokenWithQuery(accessKey, secretKey, query)

	return RequestToServer(reqURL, "GET", tokenString, query)
}

// DeleteOrder 주문 취소 접수 - 주문 UUID를 통해 해당 주문에 대한 취소 접수
func DeleteOrder(accessKey string, secretKey string, query map[string]interface{}) []byte {

	reqURL := baseURL + "/v1/order"
	tokenString := jwt.GetJwtTokenWithQuery(accessKey, secretKey, query)

	return RequestToServer(reqURL, "DELETE", tokenString, query)
}

// PostOrder 주문하기
func PostOrder(accessKey string, secretKey string, query map[string]interface{}) []byte {

	reqURL := baseURL + "/v1/orders"
	tokenString := jwt.GetJwtTokenWithQuery(accessKey, secretKey, query)

	return RequestToServer(reqURL, "POST", tokenString, query)
}

// GetMarketCode 마켓 코드 조회 - 업비트에서 거래 가능한 마켓 목록
func GetMarketCode() []byte {

	reqURL := baseURL + "/v1/market/all"

	client := &http.Client{}
	req, err := http.NewRequest("GET", reqURL, nil)
	myerr.CheckErr(err)
	res, err := client.Do(req)
	myerr.CheckErr(err)
	bytes, err := ioutil.ReadAll(res.Body)
	myerr.CheckErr(err)
	defer res.Body.Close()

	return bytes
}

// RequestToServer 업비트 서버로 요청
func RequestToServer(reqURL string, method string, tokenString string, query map[string]interface{}) []byte {

	client := &http.Client{}
	req, err := http.NewRequest(method, reqURL, nil)
	myerr.CheckErr(err)

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
	req.URL.RawQuery = q.Encode()

	fmt.Println(q.Encode())

	req.Header.Add("Authorization", "Bearer "+tokenString)
	res, err := client.Do(req)
	myerr.CheckErr(err)

	bytes, err := ioutil.ReadAll(res.Body)
	myerr.CheckErr(err)

	defer res.Body.Close()

	return bytes
}

// ConvertStructToMap struct -> map[string]interface{}
func ConvertStructToMap(object interface{}) map[string]interface{} {
	conv := make(map[string]interface{})
	tmp, err := json.Marshal(object)
	myerr.CheckErr(err)
	json.Unmarshal(tmp, &conv)
	return conv
}
