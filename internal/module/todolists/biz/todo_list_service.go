package biz

import (
	"Practice.Golang/internal/module/todolists/model"
	"Practice.Golang/internal/module/todolists/storage"
	"github.com/labstack/echo/v4"
)

type ITodoListService interface{}

type TodoListService struct {
	todoListRepo storage.ITodoListRepo
}

func NewTodoListService() *TodoListService {
	return &TodoListService{
		todoListRepo: storage.NewTodoListRepo(),
	}
}

func (s *TodoListService) GetTodoList(ctx echo.Context) *[]model.TodoList {
	return s.todoListRepo.GetAll(ctx)
}
