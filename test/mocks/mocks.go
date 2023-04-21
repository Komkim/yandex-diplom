package mocks

//go:generate mockgen -source=../../storage/repository/orders.go -destination=./packages/storagemocks/orders_mocks.go -package=storagemocks
//go:generate mockgen -source=../../storage/repository/balance.go -destination=./packages/storagemocks/balance_mocks.go -package=storagemocks
//go:generate mockgen -source=../../storage/repository/users.go -destination=./packages/storagemocks/users_mocks.go -package=storagemocks
