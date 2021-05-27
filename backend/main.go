package main

import (
	"encoding/json"
	"fmt"

	"github.com/noirstar/autotrading/backend/apis/restapi"
	"github.com/noirstar/autotrading/backend/models"
)

func main() {

	//accessKey := env.GetEnv("UPBIT_ACCESS_KEY")
	//secretKey := env.GetEnv("UPBIT_SECRET_KEY")

	req := models.ReqMinuteCandles{
		Market: "KRW-BTC",
		Count:  1,
	}
	res := []models.ResMinuteCandles{}
	reqText := restapi.GetMinuteCandles(&req, 1)
	json.Unmarshal(reqText, &res)

	fmt.Printf("%+v", res)
}
