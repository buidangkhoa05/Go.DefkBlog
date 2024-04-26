package handler

import (
	"Practice.Golang/internal/module/todolists/biz"
	"github.com/labstack/echo/v4"
	"net/http"
)

func HandleGetTodoLists(ctx echo.Context) error {
	todoListService, err := biz.NewTodoListService()

	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	todoLists := todoListService.GetTodoList(ctx)

	if todoLists == nil {
		return ctx.NoContent(http.StatusNoContent)
	}

	return ctx.JSON(http.StatusOK, todoLists)
}
