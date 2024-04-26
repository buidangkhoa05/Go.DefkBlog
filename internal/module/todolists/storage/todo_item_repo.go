package storage

import (
	"Practice.Golang/internal/module/todolists/model"
	"Practice.Golang/pkg/storage"
	"gorm.io/gorm"
)

type ITodoItemRepo interface {
	GetAll() []model.TodoItem
}

type TodoItemRepo struct {
	db *gorm.DB
}

func NewTodoItemRepo() (*TodoItemRepo, error) {
	db, err := storage.NewPostgresDb()

	if err != nil {
		return nil, err
	}

	return &TodoItemRepo{
		db: db,
	}, nil
}

func (r *TodoItemRepo) GetAll() []model.TodoItem {
	var todoItems []model.TodoItem
	r.db.Model(&model.TodoItem{}).Scan(todoItems)
	return todoItems
}
