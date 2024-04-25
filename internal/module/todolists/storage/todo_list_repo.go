package storage

import (
	"Practice.Golang/internal/module/todolists/model"
	"Practice.Golang/pkg/storage"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type ITodoListRepo interface {
	GetAll(ctx echo.Context) *[]model.TodoList
}

type TodoListRepo struct {
	Db *gorm.DB
}

func NewTodoListRepo() *TodoListRepo {
	pg, err := storage.NewPostgresDb()

	if err != nil {
		panic(err)
	}

	return &TodoListRepo{
		Db: pg,
	}
}

func (repo *TodoListRepo) GetAll(ctx echo.Context) *[]model.TodoList {
	var todoLists []model.TodoList
	repo.Db.Model(&model.TodoList{}).Select("*").Scan(&todoLists)
	return &todoLists
}
