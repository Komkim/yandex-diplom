package entity

import (
	"github.com/google/uuid"
	"time"
)

type Balance struct {
	Id       uuid.UUID `db:"id"`
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
