package app

import (
	v1 "Go.DefkBlog/internal/controller/http/v1"
	"Go.DefkBlog/pkg/server"
	"fmt"
	"github.com/labstack/echo/v4"
)

func Run() {
	var err error = nil

	e := echo.New()

	//mapping router
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
