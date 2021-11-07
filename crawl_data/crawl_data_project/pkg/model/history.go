package model

import uuid "github.com/google/uuid"

type History struct {
	ID  *uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()" json:"id,omitempty" bson:"_id,omitempty"`
	Url string     `gorm:"type:varchar(255)" json:"url"`

	//Books    *[]Book            `json:"books,omitempty"`
	BaseModel
}
