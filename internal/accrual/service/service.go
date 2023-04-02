package service

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
)

type Services struct {
	GoodsService
	AccrualOrdersService
	ConnectService
	log *zerolog.Event
}

func NewSerivces(log *zerolog.Event, db *pgxpool.Pool) *Services {
	gs := NewGoodsService(db)
	cs := NewConnectService(db)
	os := NewAccrualOrdersService(db, gs.GoodsRepo, cs.ConnectRepo)

	return &Services{
		GoodsService:         gs,
		AccrualOrdersService: os,
		ConnectService:       cs,
		log:                  log,
	}
}
