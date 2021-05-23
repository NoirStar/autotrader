package main

import (
	"fmt"
	"reflect"

	"github.com/noirstar/autotrading/backend/models"
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
	v := reflect.ValueOf(a)
	fmt.Println(v)
}
