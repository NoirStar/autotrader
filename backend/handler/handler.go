package handler

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
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

// PostRegisterUser 회원가입 핸들러
func PostRegisterUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		params := make(map[string]string)
		err := c.Bind(&params)
		if err != nil {
			fmt.Println(err)
		}
		return c.JSON(http.StatusOK, params["username"])
	}
}

// PostLogin 로그인 핸들러
func PostLogin() echo.HandlerFunc {
	return func(c echo.Context) error {
		params := make(map[string]string)
		err := c.Bind(&params)
		if err != nil {
			fmt.Println(err)
		}
		return c.JSON(http.StatusOK, params["username"])
	}
}
