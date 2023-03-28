package entity

import (
	"github.com/google/uuid"
	"time"
)

type Users struct {
	Id             uuid.UUID `db:"id"`
	Login          string    `db:"login"`
	HashedPassword string    `db:"hashed_password"`
	CreateAt       time.Time `db:"create_at"`
}

func (u *Users) BeforeSave() error {
	return nil
}

func (u *Users) Prepare() error {
	return nil
}

func (u *Users) Validate() error {
	return nil
}
