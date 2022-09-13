package account

import "finalProject/wallet"

type Account struct {
	Id       int                     `json"id"`
	UserName string                  `json:"username"`
	Password string                  `json:"password"`
	PinCode  string                  `json:"pincode"`
	Wallet   wallet.MockWalletReturn `json:"account"`
}

type Response struct {
	Id int `json:"id"`
}
