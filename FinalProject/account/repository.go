package account

import (
	"finalProject/wallet"
)

type Repository interface {
	Migrate() error
	Create(account Account, wallet wallet.MockWalletReturn) (*Account, error)
}
