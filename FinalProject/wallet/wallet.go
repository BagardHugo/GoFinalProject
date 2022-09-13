package wallet

type MockWalletCall struct {
	Blockchain string `json:"blockchain"`
	Pincode    string `json:"pin_code"`
}

type MockWalletReturn struct {
	Id              int    `json:"id"`
	WalletAddress   string `json:"wallet_address"`
	CurrencyCode    string `json:"currency_code"`
	CurrencyBalance string `json:"currency_balance"`
}

type Wallet struct {
	Id      int    `json"id"`
	Address string `json:"address"`
}
