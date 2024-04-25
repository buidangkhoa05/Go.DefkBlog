package todolists

import "Practice.Golang/internal/module/todolists/storage"

type Service struct {
	todoListRepo *storage.TodoListRepo
}

func NewService() *Service {
	return &Service{
		todoListRepo: &storage.TodoListRepo{},
	}
}
