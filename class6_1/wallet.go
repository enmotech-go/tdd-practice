package class6_1

type Wallet struct {
	balance int
}

func (wallet *Wallet) Deposit(amount int) {
	wallet.balance += amount
}

func (wallet *Wallet) GetBalance() int {
	return wallet.balance
}
