package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/noirstar/autotrader/api"
	"github.com/noirstar/autotrader/db"
	"github.com/noirstar/autotrader/model"
	"github.com/noirstar/autotrader/utils"
	"golang.org/x/crypto/bcrypt"
)

// GetIndex Return index.html
func GetIndex() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.File("public/views/index.html")
	}
}

// GetCandles Return Candles
func GetCandles() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Get Candles")
	}
}

// GetCoinInfo Return CoinInfo
func GetCoinInfo() echo.HandlerFunc {
	return func(c echo.Context) error {
		info := []model.ResMarketCode{}
		err := json.Unmarshal(api.GetMarketCode(), &info)
		if err != nil {
			fmt.Println(err)
		}
		return c.JSON(http.StatusOK, info)
	}
}

// GetMarketInfo Return Market aggregate info
func GetMarketInfo() echo.HandlerFunc {
	return func(c echo.Context) error {
		params := make(map[string]int)
		err := c.Bind(&params)
		if err != nil {
			fmt.Println(err)
			return c.String(http.StatusInternalServerError, "파라미터 바인딩에 실패했습니다")
		}

		if len(params) != 1 {
			return c.String(http.StatusBadRequest, "잘못된 요청입니다")
		}

		market, err := db.FindMarketData(time.Duration(params["min"]))
		if err != nil {
			fmt.Println(err)
			return c.String(http.StatusInternalServerError, "마켓데이터 요청에 실패했습니다")
		}

		data, err := json.Marshal(market)

		return c.JSON(http.StatusOK, data)
	}
}

// PostRegisterUser 회원가입 핸들러
func PostRegisterUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		params := make(map[string]string)
		err := c.Bind(&params)
		if err != nil {
			fmt.Println(err)
			return c.String(http.StatusInternalServerError, "파라미터 바인딩에 실패했습니다")
		}

		if len(params) != 5 {
			return c.String(http.StatusBadRequest, "잘못된 요청입니다")
		}

		pwHash, err := bcrypt.GenerateFromPassword([]byte(params["pw"]), bcrypt.DefaultCost)
		if err != nil {
			fmt.Println(err)
			return c.String(http.StatusInternalServerError, "패스워드 해싱에 실패했습니다")
		}

		user := &model.User{
			ID:       params["id"],
			PW:       string(pwHash),
			Email:    params["email"],
			Nickname: params["nickname"],
			Birth:    params["birth"],
			Level:    1,
			Money:    10000000,
		}
		err = db.CreateUser(user)
		if err != nil {
			fmt.Println(err)
			return c.String(http.StatusInternalServerError, "회원가입에 실패했습니다")
		}
		return c.String(http.StatusOK, params["nickname"])
	}
}

// PostLogin 로그인 핸들러
func PostLogin() echo.HandlerFunc {
	return func(c echo.Context) error {
		params := make(map[string]string)
		err := c.Bind(&params)
		if err != nil {
			fmt.Println(err)
			return c.String(http.StatusInternalServerError, "파라미터 바인딩에 실패했습니다")
		}

		if len(params) != 2 {
			return c.String(http.StatusBadRequest, "잘못된 요청입니다")
		}

		user, err := db.LoginUser(params["id"], params["pw"])

		if err != nil {
			fmt.Println(err)
			return c.String(http.StatusInternalServerError, "아이디 혹은 패스워드가 일치하지 않습니다")
		}

		// claim 생성
		claims := map[string]interface{}{
			"nickname": user.Nickname,
		}

		accessToken, err := utils.GenerateJwt(c, claims)
		if err != nil {
			fmt.Println(err)
			return c.String(http.StatusInternalServerError, "JWT 토큰 발급에 실패하였습니다")
		}
		fmt.Println(user.Nickname, ":", accessToken)

		return c.String(http.StatusOK, user.Nickname)
	}
}

// PostCheck 중복값 검사 핸들러
func PostCheck() echo.HandlerFunc {
	return func(c echo.Context) error {
		params := make(map[string]string)
		err := c.Bind(&params)
		if err != nil {
			fmt.Println(err)
			return c.String(http.StatusInternalServerError, "파라미터 바인딩에 실패했습니다")
		}
		if len(params) > 1 {
			return c.String(http.StatusBadRequest, "잘못된 요청입니다")
		}

		check, err := db.CheckDuplicate(params)

		if check {
			return c.String(http.StatusOK, "true")
		}
		return c.String(http.StatusOK, "false")
	}
}
