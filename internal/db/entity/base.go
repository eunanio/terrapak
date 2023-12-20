package entity

import (
	"time"

	"github.com/google/uuid"
)

type ModelBase struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid();" json:"id"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}