package application

type Services struct {
	BalanceService BalanceService
	OrdersService  OrdersService
	UsersService   UsersService
}

func NewServices() Services {
	return Services{
		BalanceService: *NewBalanceService(),
		OrdersService:  *NewOrdersService(),
		UsersService:   *NewUsersService(),
	}
}
