package model

import (
	"time"

	uuid "github.com/google/uuid"
	"gorm.io/gorm"
)

type Pagination struct {
	Page     int
	PageSize int
}

type UriParse struct {
	ID []string `json:"id" uri:"id"`
}

// MsMetadata describes the structure.
type BaseModel struct {
	CreatorID *uuid.UUID      `json:"creator_id,omitempty" gorm:"type:uuid;;default:uuid_generate_v4()"`
	UpdaterID *uuid.UUID      `json:"updater_id,omitempty" gorm:"type:uuid;;default:uuid_generate_v4()"`
	CreatedAt time.Time       `gorm:"column:created_at;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time       `gorm:"column:updated_at;default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt *gorm.DeletedAt `json:"deleted_at,omitempty"`
}

type BaseModelStripped struct {
	ID uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()" json:"id"`
}
