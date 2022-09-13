package account

import "finalProject/wallet"

type Account struct {
	Id       int64         `json"id"`
	UserName string        `json:"username"`
	Password string        `json:"password"`
	PinCode  string        `json:"pincode"`
	Wallet   wallet.Wallet `json:"account"`
}

type Response struct {
	Id int `json:"id"`
}
