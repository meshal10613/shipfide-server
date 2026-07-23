package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// UUIDBase can be embedded in models that want automatic UUID generation.
// Note: models that define their own BeforeCreate do NOT need to embed this.
type UUIDBase struct {
	ID string `gorm:"type:uuid;primaryKey"`
}

func (b *UUIDBase) BeforeCreate(tx *gorm.DB) error {
	if b.ID == "" {
		b.ID = uuid.NewString()
	}
	return nil
}
