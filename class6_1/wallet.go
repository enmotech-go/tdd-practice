package class6_1

type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

func (wallet *Wallet) Deposit(amount Bitcoin) {
	wallet.balance += amount
}

func (wallet *Wallet) GetBalance() Bitcoin {
	return wallet.balance
}
