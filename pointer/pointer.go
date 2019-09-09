package pointer

import (
	"errors"
	"fmt"
)

type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

var InsufficientFunds = errors.New("insufficient funds")

func (w *Wallet) Withdraw(coin Bitcoin) error {
	if coin > w.balance {
		return InsufficientFunds
	}
	w.balance -= coin
	return nil
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
