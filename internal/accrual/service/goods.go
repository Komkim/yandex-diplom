package service

import (
	"yandex-diplom/internal/accrual/model"
	storage "yandex-diplom/storage/repository"

	"github.com/jackc/pgx/v5/pgxpool"
)

type GoodsService struct {
	GoodsRepo model.GoodsRepo
}

func NewGoodsService(db *pgxpool.Pool) GoodsService {
	return GoodsService{model.NewGoods(db)}
}

func (g *GoodsService) SetGoods(goods *storage.Goods) error {
	_, err := g.GoodsRepo.SetGoods(goods.Description, goods.Price)
	return err
}

func (g *GoodsService) SetReward(reward *storage.Reward) error {
	//goods, err := g.GoodsRepo.GetGoods(reward.Match)
	//if err != nil {
	//	return err
	//}

	return g.GoodsRepo.SetReward(reward.Match, reward.RewardType, reward.Reward)
}
