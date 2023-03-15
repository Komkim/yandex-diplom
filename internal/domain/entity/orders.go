package entity

import (
	"github.com/google/uuid"
	"time"
)

type status string

func (s status) Status() status {
	return s
}

const (
	REGISTERED status = "REGISTERED"
	INVALID           = "INVALID"
	PROCESSING        = "PROCESSING"
	PROCESSED         = "PROCESSED"
)

type Orders struct {
	Id        uuid.UUID `db:"id"`
	UserId    uuid.UUID `db:"user_id"`
	BalanceId uuid.UUID `db:"balance_id"`
	Number    int64     `db:"number"`
	Status    status    `db:"status"`
	Accrual   float64   `db:"accrual"`
	Withdraw  float64   `db:"withdraw"`
	CreateAt  time.Time `db:"create_at"`
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
