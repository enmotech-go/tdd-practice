package pointer

type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(coin Bitcoin) {
	w.balance += coin
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
