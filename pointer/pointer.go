package pointer

type Wallet struct {
	balance int
}

func (w *Wallet) Deposit(coin int) {
	w.balance += coin
}

func (w *Wallet) Balance() int {
	return w.balance
}
