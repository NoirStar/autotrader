package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/noirstar/autotrading/backend/models"
	"github.com/noirstar/autotrading/backend/utils/env"
	"github.com/noirstar/autotrading/backend/utils/myerr"
)

func main() {

}

func test() {
	accessKey := env.GetEnv("UPBIT_ACCESS_KEY")
	secretKey := env.GetEnv("UPBIT_SECRET_KEY")
	url := env.GetEnv("UPBIT_BASE_URL") + "/v1/accounts"
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
