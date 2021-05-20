package restapi

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/noirstar/autotrading/backend/models"
	"github.com/noirstar/autotrading/backend/utils/env"
	"github.com/noirstar/autotrading/backend/utils/myerr"
)

var baseURL = env.GetEnv("UPBIT_BASE_URL")

// GetAccount 전체 계좌 조회
func GetAccount(accessKey string, secretKey string) {

	url := baseURL + "/v1/accounts"

	tokenString := models.GetJwtToken(accessKey, secretKey)
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	myerr.CheckErr(err)
	req.Header.Add("Authorization", "Bearer "+tokenString)
	res, err := client.Do(req)
	myerr.CheckErr(err)

	defer res.Body.Close()

	bytes, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(bytes))
}

// GetOrderChance 주문 가능 정보 - 마켓별 주문 가능 정보 확인
func GetOrderChance(accessKey string, secretKey string, query map[string]string) {

	url := baseURL + "/v1/orders/chance"

	tokenString := models.GetJwtTokenWithQuery(accessKey, secretKey, query)
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	myerr.CheckErr(err)
	req.Header.Add("Authorization", "Bearer "+tokenString)
	res, err := client.Do(req)
	myerr.CheckErr(err)

	defer res.Body.Close()

	bytes, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(bytes))
}
