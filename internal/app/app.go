package app

import (
	v1 "Practice.Golang/internal/controller/http/v1"
	"Practice.Golang/pkg/server"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func Run() {
	var err error = nil

	e := echo.New()

	//Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())
	e.Use(middleware.Recover())

	//Router
	//Router for swagger
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	v1.MapRouter(e)

	apiServer := server.NewApiServer(e)

	select {
	case err = <-apiServer.Notify():
		fmt.Println("app - Run - apiServer.Notify()")
	}

	err = apiServer.Shutdown()
	if err != nil {
		fmt.Println("app - Run - apiServer.Shutdown()")
	}
}
