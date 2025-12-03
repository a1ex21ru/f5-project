package models

import (
	"gorm.io/gorm"
	"time"
)

type Note struct {
	gorm.Model
	Name        string `json:"name" gorm:"column:name"`
	Description string `json:"description,omitempty" gorm:"column:description"`
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
