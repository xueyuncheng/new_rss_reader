package model

import (
	"time"

	"gorm.io/gorm"
)

type Model struct {
	ID        int            `json:"id"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-"`
}
