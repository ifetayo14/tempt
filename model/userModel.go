package model

import (
	"github.com/google/uuid"
	"time"
)

type (
	User struct {
		Id        uuid.UUID `json:"id" gorm:"column:id"`
		Username  string    `json:"username" gorm:"column:username"`
		Password  string    `json:"password" gorm:"column:password"`
		CreatedAt time.Time `json:"created_at" gorm:"column:created_at"`
		UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at"`
		DeletedAt time.Time `json:"deleted_at" gorm:"column:deleted_at"`
	}

	RegisterRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	RegisterResponse struct {
		Token string `json:"token"`
	}
)

func (User) TableName() string {
	return "user"
}
