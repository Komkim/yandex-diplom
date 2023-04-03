package updateorder

import (
	"context"
	"github.com/rs/zerolog"
	"strconv"
	"time"
	"yandex-diplom/config"
	storage "yandex-diplom/storage/repository"
)

const (
	POLL = 5
)

type AccrualOrder struct {
	Order   string  `json:"order"`
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

func (a *Accrual) Start(ctx context.Context) {
	orderChan := make(chan []storage.Order)
	go a.basePoll(ctx, orderChan)
	go a.sendAccrual(ctx, orderChan)
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

func (a *Accrual) sendAccrual(ctx context.Context, orderChan chan []storage.Order) {

	for {
		select {
		case orders := <-orderChan:
			resultOrders := make([]storage.OrderAccrual, 0, len(orders))
		n:
			for _, o := range orders {
				n, err := strconv.ParseInt(o.Number, 10, 64)
				if err != nil {
					a.log.Err(err)
					continue n
				}
				order, err := a.MyClient.GetAccrual(n)
				if err != nil {
					a.log.Err(err)
					continue n
				}
				if order == nil {
					continue n
				}
				resultOrders = append(resultOrders, *order)
			}

			go func() {
				err := a.storage.SetAccrual(resultOrders)
				a.log.Err(err)
			}()

		case <-ctx.Done():
			return
		}
	}
}
