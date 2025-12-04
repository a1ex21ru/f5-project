package models

import (
	"gorm.io/gorm"
	"time"
)

type Note struct {
	Name        string `json:"name" gorm:"column:name"`
	Description string `json:"description,omitempty" gorm:"column:description"`
	gorm.Model
}

type Notes struct {
	Notes []Note `json:"data"`
}

func (Note) TableName() string {
	return "notes"
}

type UpdateNote struct {
	Name        string `json:"name" gorm:"column:name"`
	Description string `json:"description,omitempty" gorm:"column:description"`
	UpdatedAt   time.Time
}

func (UpdateNote) TableName() string {
	return "notes"
}

type NoteResponse struct {
	Message string `json:"message"`
	Data    Note   `json:"data,omitempty"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}
