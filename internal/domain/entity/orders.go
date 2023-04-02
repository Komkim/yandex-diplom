package entity

import (
	"github.com/google/uuid"
	"time"
)

type Orders struct {
	ID       uuid.UUID `db:"id"`
	UserID   uuid.UUID `db:"user_id"`
	Number   int64     `db:"number"`
	Status   string    `db:"status"`
	Sum      *float64  `db:"sum"`
	CreateAt time.Time `db:"create_at"`
}

func (o *Orders) BeforeSave() error {
	return nil
}

func (o *Orders) Prepare() error {
	return nil
}

func (o *Orders) Validate() error {
	return nil
}
