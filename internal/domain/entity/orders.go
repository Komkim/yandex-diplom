package entity

import (
	"github.com/google/uuid"
	"time"
)

type status string

func (s status) Status() status {
	return s
}

func (s status) String() string {
	return s.String()
}

const (
	REGISTERED status = "REGISTERED"
	INVALID           = "INVALID"
	PROCESSING        = "PROCESSING"
	PROCESSED         = "PROCESSED"
)

type Orders struct {
	Id       uuid.UUID `db:"id"`
	UserId   uuid.UUID `db:"user_id"`
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
