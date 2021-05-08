package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/noirstar/autotrading/backend/utils/env"
	"github.com/noirstar/autotrading/backend/utils/myerr"
)

func main() {
	url := env.GetEnv("UPBIT_BASE_URL") + "/v1/accounts"
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"access_key": env.GetEnv("UPBIT_ACCESS_KEY"),
		"nonce":      uuid.New(),
	})
	tokenString, err := token.SignedString([]byte(env.GetEnv("UPBIT_SECRET_KEY")))
	myerr.CheckErr(err)

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
