package entity

import (
	"github.com/google/uuid"
	"time"
)

type Balance struct {
	Id       uuid.UUID `db:"id"`
	User_id  uuid.UUID `db:"user_id"`
	Number   int64     `db:"number"`
	Current  float64   `db:"current"`
	Withdraw float64   `db:"withdraw"`
	UploadAt time.Time `db:"upload_at"`
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
