package model

import (
	"gorm.io/gorm"
	"time"
)

type BaseModel struct {
	ID        uint64    `json:"id" gorm:"column:id"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
}

type SoftDeleteModel struct {
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
}
