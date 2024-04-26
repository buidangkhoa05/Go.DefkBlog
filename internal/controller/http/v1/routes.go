package v1

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type GroupRoute *echo.Group

func MapRouter(e *echo.Echo) {

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World!")
	})

	apiGp := e.Group("/api")
	ver1Gp := apiGp.Group("/v1")
	{
		// Test endpoint
		ver1Gp.GET("/hello", func(c echo.Context) error {
			return c.String(http.StatusOK, "Hello word")
		})

		//TodoList
		UseTodoListRoute(ver1Gp)
	}

}
