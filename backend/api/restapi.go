package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"

	"github.com/noirstar/autotrader/model"
	"github.com/noirstar/autotrader/utils"
)

var baseURL = utils.GetEnv("UPBIT_BASE_URL")

// GetAccount 전체 계좌 조회
func GetAccount(accessKey string, secretKey string) []byte {

	reqURL := baseURL + "/v1/accounts"
	tokenString, err := utils.CreateUpbitJwt(accessKey, secretKey)
	utils.CheckErr(err)

	return RequestToServer(reqURL, "GET", tokenString, nil)
}

// GetOrderChance 주문 가능 정보 - 마켓별 주문 가능 정보 확인
func GetOrderChance(accessKey string, secretKey string, query *model.ReqChance) []byte {

	reqURL := baseURL + "/v1/orders/chance"
	queryMap := ConvertStructToMap(query)
	tokenString, err := utils.CreateUpbitJwtQuery(accessKey, secretKey, queryMap)
	utils.CheckErr(err)

	return RequestToServer(reqURL, "GET", tokenString, queryMap)
}

// GetOrderSearch 개별 주문 조회 - 주문 UUID 를 통해 개별 주문건을 조회
func GetOrderSearch(accessKey string, secretKey string, query *model.ReqOrderSearch) []byte {

	reqURL := baseURL + "/v1/order"
	queryMap := ConvertStructToMap(query)
	tokenString, err := utils.CreateUpbitJwtQuery(accessKey, secretKey, queryMap)
	utils.CheckErr(err)

	return RequestToServer(reqURL, "GET", tokenString, queryMap)
}

// GetOrdersSearch 주문 리스트 조회 - 주문 리스트를 조회
func GetOrdersSearch(accessKey string, secretKey string, query *model.ReqOrdersSearch) []byte {

	reqURL := baseURL + "/v1/orders"
	queryMap := ConvertStructToMap(query)
	tokenString, err := utils.CreateUpbitJwtQuery(accessKey, secretKey, queryMap)
	utils.CheckErr(err)

	return RequestToServer(reqURL, "GET", tokenString, queryMap)
}

// DeleteOrder 주문 취소 접수 - 주문 UUID를 통해 해당 주문에 대한 취소 접수
func DeleteOrder(accessKey string, secretKey string, query *model.ReqDeleteOrder) []byte {

	reqURL := baseURL + "/v1/order"
	queryMap := ConvertStructToMap(query)
	tokenString, err := utils.CreateUpbitJwtQuery(accessKey, secretKey, queryMap)
	utils.CheckErr(err)

	return RequestToServer(reqURL, "DELETE", tokenString, queryMap)
}

// PostOrder 주문하기
func PostOrder(accessKey string, secretKey string, query *model.ReqOrders) []byte {

	reqURL := baseURL + "/v1/orders"
	queryMap := ConvertStructToMap(query)
	tokenString, err := utils.CreateUpbitJwtQuery(accessKey, secretKey, queryMap)
	utils.CheckErr(err)

	return RequestToServer(reqURL, "POST", tokenString, queryMap)
}

// GetMarketCode 마켓 코드 조회 - 업비트에서 거래 가능한 마켓 목록
func GetMarketCode() []byte {

	reqURL := baseURL + "/v1/market/all?isDetails=true"

	return RequestToServerSimple(reqURL, "GET", nil)
}

// GetMinuteCandles 분 캔들 조회
func GetMinuteCandles(query *model.ReqMinuteCandles, unit int) []byte {

	reqURL := baseURL + "/v1/candles/minutes/" + strconv.Itoa(unit)

	return RequestToServerSimple(reqURL, "GET", ConvertStructToMap(query))
}

// GetDayCandles 일 캔들 조회
func GetDayCandles(query *model.ReqDayCandles) []byte {

	reqURL := baseURL + "/v1/candles/days"

	return RequestToServerSimple(reqURL, "GET", ConvertStructToMap(query))
}

// GetWeekCandles 주 캔들 조회
func GetWeekCandles(query *model.ReqWeekCandles) []byte {

	reqURL := baseURL + "/v1/candles/weeks"

	return RequestToServerSimple(reqURL, "GET", ConvertStructToMap(query))
}

// GetMonthsCandles 달 캔들 조회
func GetMonthsCandles(query *model.ReqMonthCandles) []byte {

	reqURL := baseURL + "/v1/candles/months"

	return RequestToServerSimple(reqURL, "GET", ConvertStructToMap(query))
}

// RequestToServer 업비트 서버로 요청
func RequestToServer(reqURL string, method string, tokenString string, query map[string]interface{}) []byte {

	client := &http.Client{}
	req, err := http.NewRequest(method, reqURL, nil)
	utils.CheckErr(err)

	q := url.Values{}

	for key, value := range query {
		switch val := value.(type) {
		case string:
			q.Add(key, value.(string))
		case int, uint32, uint64, float64:
			q.Add(key, fmt.Sprint(value))
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

	//fmt.Println(q.Encode())

	req.Header.Add("Authorization", "Bearer "+tokenString)
	res, err := client.Do(req)
	utils.CheckErr(err)
	bytes, err := ioutil.ReadAll(res.Body)
	utils.CheckErr(err)
	defer res.Body.Close()

	return bytes
}

// RequestToServerSimple 토큰 미포함 요청
func RequestToServerSimple(reqURL string, method string, query map[string]interface{}) []byte {

	client := &http.Client{}
	req, err := http.NewRequest(method, reqURL, nil)
	utils.CheckErr(err)

	q := req.URL.Query()

	for key, value := range query {
		switch val := value.(type) {
		case string:
			q.Add(key, value.(string))
		case int, uint32, uint64, float64:
			q.Add(key, fmt.Sprint(value))
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

	//fmt.Println(q.Encode())

	res, err := client.Do(req)
	utils.CheckErr(err)
	bytes, err := ioutil.ReadAll(res.Body)
	utils.CheckErr(err)
	defer res.Body.Close()

	return bytes

}

// ConvertStructToMap struct -> map[string]interface{}
// 데이터를 unmarshal 하는 과정중 숫자 데이터는 float64로 변한됨. 큰값 전송시, 오버플로 현상 발생가능 (이러면 모델 숫자 타입을 정의하는 의미가.. 없는데)
func ConvertStructToMap(object interface{}) map[string]interface{} {
	conv := make(map[string]interface{})
	tmp, err := json.Marshal(object)
	utils.CheckErr(err)
	json.Unmarshal(tmp, &conv)
	return conv
}
