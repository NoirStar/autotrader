package main

import (
	"encoding/json"
	"fmt"

	"github.com/noirstar/autotrading/backend/models"
	"github.com/noirstar/autotrading/backend/utils/myerr"
)

func main() {

	// accessKey := env.GetEnv("UPBIT_ACCESS_KEY")
	// secretKey := env.GetEnv("UPBIT_SECRET_KEY")

	// a := make(map[string]string)
	// a["market"] = "KRW-BTC"
	// res := restapi.GetOrderChance(accessKey, secretKey, a)
	// b := models.ResChance{}

	// err := json.Unmarshal(res, &b)
	// myerr.CheckErr(err)
	// fmt.Println(b)
	// fmt.Println(string(res))
	a := models.ReqOrders{}
	a.Identifier = "ab"
	e, err := json.Marshal(a)
	myerr.CheckErr(err)
	fmt.Println(string(e))
}
