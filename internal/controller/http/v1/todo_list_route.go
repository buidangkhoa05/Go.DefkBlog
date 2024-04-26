package v1

import (
	"Practice.Golang/internal/module/todolists/handler"
	"github.com/labstack/echo/v4"
)

func UseTodoListRoute(gp *echo.Group) {
	gp.GET("/todo-lists", handler.HandleGetTodoLists)
}
