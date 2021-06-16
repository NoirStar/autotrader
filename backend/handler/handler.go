package handler

import (
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
