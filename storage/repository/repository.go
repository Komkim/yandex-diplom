package storage

type Storage struct {
	Orders  Orders
	Balance Balance
	Users   Users
}

type St interface {
	Orders
	Balance
	Users
}

type Stt struct {
	Orders
	Balance
	Users
}

func NewStorage(orders Orders, balance Balance, users Users) Storage {
	return Storage{
		Orders:  orders,
		Balance: balance,
		Users:   users,
	}
}
