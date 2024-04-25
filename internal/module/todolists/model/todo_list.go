package model

import "time"

type TodoList struct {
	TodoLstId   int        `json:"todo_lst_id" gorm:"column:todo_lst_id"`
	Task        string     `json:"task" gorm:"column:task"`
	IsCompleted bool       `json:"is_completed" gorm:"column:is_completed"`
	CreatedDate *time.Time `json:"created_date" gorm:"column:created_date"`
	UpdatedDate *time.Time `json:"updated_date" gorm:"column:updated_date"`
	CreatedBy   int        `json:"created_by" gorm:"column:created_by"`
	Updated     int        `json:"updated_by" gorm:"column:updated"`
	IsDeleted   bool       `json:"is_deleted" gorm:"column:is_deleted"`
}

func (TodoList) TableName() string {
	return "todo_list"
}
