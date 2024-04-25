package todolists

import (
	"Practice.Golang/internal/module/todolists/biz"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	todoListService *biz.TodoListService
}

func (h Handler) HandleGetTodoLst(ctx echo.Context) error {
	h.todoListService.GetTodoList(ctx)
	return nil
}
