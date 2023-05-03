package entity

import (
	"time"

	"github.com/google/uuid"
)

type Balance struct {
	ID       uuid.UUID `db:"id"`
	UserID   uuid.UUID `db:"user_id"`
	Sum      float64   `db:"sum"`
	CreateAt time.Time `db:"create_at"`
}

func (b *Balance) BeforeSave() error {
	return nil
}

func (b *Balance) Prepare() error {
	return nil
}

func (b *Balance) Validate() error {
	return nil
}
