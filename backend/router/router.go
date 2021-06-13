package router

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/noirstar/autotrader/handler"
)

// New Initalize Webserver
func New() *echo.Echo {

	const (
		indexPath   string = "public/views/index.html"
		faviconPath string = "public/favicon.ico"
		staticPath  string = "public/static"
	)

	e := echo.New()
	e.Logger.SetLevel(log.DEBUG)
	e.Use(middleware.Logger())
	e.Static("/static", staticPath)

	e.GET("/", handler.GetIndex())

	v1 := e.Group("/api/v1")
	v1.GET("/candles", handler.GetCandles())

	return e
}
