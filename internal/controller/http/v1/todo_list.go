package v1

import (
	"Practice.Golang/internal/module/todolists"
	"github.com/labstack/echo/v4"
)

func AddTodoListRoute(gp *echo.Group) {
	handler := &todolists.Handler{}

	gp.GET("/todo-lists", handler.HandleGetTodoLst)
}
