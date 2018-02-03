package blockchain

func (cli *CLI) CreateWallet() (string) {
	wallets, _ := NewWallets()
	address := wallets.CreateWallet()
	wallets.SaveToFile()

	return address
}
