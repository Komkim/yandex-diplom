package accrual

import (
	"context"
	"github.com/rs/zerolog"
	"time"
	"yandex-diplom/config"
	storage "yandex-diplom/storage/repository"
)

const (
	POLL = 30
)

type AccrualOrder struct {
	Order   int64   `json:"order"`
	Status  string  `json:"status"`
	Accrual float64 `json:"accrual"`
}

type Accrual struct {
	cfg      *config.Server
	log      *zerolog.Event
	storage  storage.Storage
	MyClient *MyClient
}

func NewAccrual(cfg *config.Server, log *zerolog.Event, storage storage.Storage) *Accrual {
	return &Accrual{
		cfg:      cfg,
		log:      log,
		storage:  storage,
		MyClient: NewClient(cfg),
	}
}

func (a *Accrual) Start() {

}

func (a *Accrual) basePoll(ctx context.Context, orderChan chan []storage.Order) {
	ticker := time.NewTicker(time.Second * POLL)

	for {
		select {
		case <-ticker.C:
			orders, err := a.storage.GetAccrualOrder()
			if err != nil {
				a.log.Err(err)
				continue
			}

			orderChan <- orders
		case <-ctx.Done():
			return
		}
	}
}

func (a *Accrual) sendAccrual(ctx context.Context, orderChan chan []storage.Order, dbChan chan []storage.Order) {

	for {
		select {
		case orders := <-orderChan:
			resultOrders := make([]storage.Order, 0, len(orders))
			orderMap := make(map[int64]storage.Order)
		n:
			for _, o := range orders {
				order, err := a.MyClient.GetAccrual(o.Number)
				if err != nil {
					a.log.Err(err)
					continue n
				}
				resultOrders = append(resultOrders, *order)
				orderMap[order.Number] = *order
			}
			dbChan <- resultOrders

			go func() {
				err := a.storage.SetAccrual(resultOrders, orderMap)
				a.log.Err(err)
			}()

		case <-ctx.Done():
			return
		}
	}
}
