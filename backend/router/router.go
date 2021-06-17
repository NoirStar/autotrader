package router

import (
	"net/http"

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

	// set middleare
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format:           "time=${time_custom} : method=${method}, uri=${uri}, status=${status}, ip=${remote_ip}\n",
		CustomTimeFormat: "2006-01-02 15:04:05",
	}))
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:8080"},
		AllowMethods: []string{http.MethodGet, http.MethodPost},
	}))

	e.Static("/static", staticPath)

	// set routes
	v1 := e.Group("/api/v1")

	e.GET("/", handler.GetIndex())
	v1.GET("/candles", handler.GetCandles())
	v1.POST("/signup", handler.PostRegisterUser())
	v1.POST("/login", handler.PostLogin())

	return e
}
