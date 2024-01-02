package models

import "github.com/google/uuid"

type Example struct {
	Id        uuid.UUID `json:"id,string" gorm:"type:uuid;default:uuid_generate_v4()"`
	Data      string    `json:"data" binding:"required"`
	CreatedAt int64     `json:"created_at,omitempty" gorm:"autoCreateTime"`
	UpdatedAt int64     `json:"updated_at,omitempty" gorm:"autoUpdateTime"`
}

func (e *Example) TableName() string {
	return "examples"
}
