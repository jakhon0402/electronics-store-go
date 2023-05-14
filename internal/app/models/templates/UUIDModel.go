package templates

import "github.com/google/uuid"

type UUIDModel struct {
	ID uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()" json:"id"`
}
