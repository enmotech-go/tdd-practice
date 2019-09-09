package pointer

import "fmt"

type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Withdraw(coin Bitcoin) {
	w.balance -= coin
}

func (w *Wallet) Deposit(coin Bitcoin) {
	w.balance += coin
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

type Stringer interface {
	String() string
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
