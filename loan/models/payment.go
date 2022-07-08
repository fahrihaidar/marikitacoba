package models

import (
	"time"

	"github.com/google/uuid"
)

type Payment struct {
	ID          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey;"`
	PaymentCode string
	Name        string
	Status      string
	Created     time.Time `gorm:"autoCreateTime"`
	Updated     time.Time `gorm:"autoUpdateTime"`
}
