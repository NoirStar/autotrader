package router

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/noirstar/autotrader/db"
	"github.com/noirstar/autotrader/handler"
	"github.com/noirstar/autotrader/utils"
)

// New Initalize Webserver
func New() *echo.Echo {

	e := echo.New()
	e.Logger.SetLevel(log.DEBUG)

	// set middleare
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format:           "time=${time_custom} : method=${method}, uri=${uri}, status=${status}, ip=${remote_ip}\n",
		CustomTimeFormat: "2006-01-02 15:04:05",
	}))
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowCredentials: true,
		AllowOrigins:     []string{"http://localhost:8080"},
		AllowMethods:     []string{http.MethodGet, http.MethodPost},
	}))

	e.Static("/static", "public/views/index.html")

	// connect db
	db.New()

	// set routes
	v1 := e.Group("/api/v1")
	v1Auth := e.Group("/api/v1")

	v1Auth.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte(utils.GetEnv("JWT_SECRET_KEY")),
		Claims:     &utils.JwtCustomClaims{},
	}))

	e.GET("/", handler.GetIndex())
	v1Auth.GET("/candles", handler.GetCandles())
	v1Auth.GET("/coins", handler.GetCoinInfo())
	v1Auth.GET("/market", handler.GetMarketInfo())
	v1.POST("/check", handler.PostCheck())
	v1.POST("/signup", handler.PostRegisterUser())
	v1.POST("/login", handler.PostLogin())

	return e
}
