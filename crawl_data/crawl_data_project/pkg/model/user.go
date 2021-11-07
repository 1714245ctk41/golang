package model

import (
	uuid "github.com/google/uuid"
)

//User represents users table in database
type User struct {
	ID       uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()" json:"id,omitempty" bson:"_id,omitempty"`
	Name     string    `gorm:"type:varchar(255)" json:"name"`
	Email    string    `gorm:"uniqueIndex;type:varchar(255)" json:"email"`
	Password string    `gorm:"->;<-;not null" json:"-"`
	Token    string    `gorm:"-" json:"token,omitempty"`
	Role     string    `gorm:"type:varchar(255)" json:"role"`
	//Books    *[]Book            `json:"books,omitempty"`
	BaseModel
}

type LoginDTO struct {
	Email    string `json:"email" form:"email" binding:"required,email"`
	Password string `json:"password" form:"password" binding:"required"`
}
type RegisterDTO struct {
	Name     string `json:"name" form:"name" binding:"required"`
	Email    string `json:"email" form:"email" binding:"required,email" `
	Password string `json:"password" form:"password" binding:"required"`
	Role     string `json:"role" form:"role" binding:"required"`
}
type UserUpdateDTO struct {
	ID       string `json:"id,omitempty" form:"id" bson:"_id,omitempty"`
	Name     string `json:"name" form:"name" binding:"required"`
	Email    string `json:"email" form:"email" binding:"required,email"`
	Password string `json:"password,omitempty" form:"password,omitempty"`
	Role     string `json:"role" form:"role" binding:"required"`
}
