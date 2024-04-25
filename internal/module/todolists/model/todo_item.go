package model

import "time"

type TodoItem struct {
	TodoItemId  string     `json:"todo_item_id" gorm:"column:todo_item_id"`
	Task        string     `json:"task" gorm:"column:task"`
	IsComplete  bool       `json:"is_complete" gorm:"column:is_complete"`
	CreatedDate *time.Time `json:"created_date" gorm:"column:created_date"`
	UpdatedDate *time.Time `json:"updated_date" gorm:"column:updated_date"`
	CreatedBy   int        `json:"created_by" gorm:"column:created_by"`
	Updated     int        `json:"updated_by" gorm:"column:updated"`
	IsDeleted   bool       `json:"is_deleted" gorm:"column:is_deleted"`
}

func (TodoItem) TableName() string {
	return "todo_item"
}
