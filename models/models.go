package models

import (
	"time"

	"gorm.io/gorm"
)

type Contact struct {
	gorm.Model
	Name  string `json:"name" gorm:"type:varchar(100);not null"`
	Tasks []Task `json:"tasks" gorm:"foreignKey:ContactID"`
}

type TaskPriority string

const (
	PriorityLow    TaskPriority = "LOW"
	PriorityMedium TaskPriority = "MEDIUM"
	PriorityHigh   TaskPriority = "HIGH"
)

type Task struct {
	gorm.Model
	Title       string       `json:"title" gorm:"type:varchar(100);not null" validate:"required,min=3,max=100"`
	Description string       `json:"description" gorm:"type:text;not null" validate:"required"`
	Priority    TaskPriority `json:"priority" gorm:"type:varchar(10);not null" validate:"required,oneof=LOW MEDIUM HIGH"`
	DueDate     time.Time    `json:"dueDate" gorm:"type:timestamp" validate:"required"`
	ContactID   uint         `json:"contactId" gorm:"not null"` 
}

