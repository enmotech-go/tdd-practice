package class6

import (
	"errors"
)

type Bitcoin int
type Wallet struct {
	balance Bitcoin
}

var InsufficientFundsError = errors.New("cannot withdraw, insufficient funds")

func (wallet *Wallet) Deposit(amount Bitcoin) {
	wallet.balance += amount
}

func (wallet *Wallet) Balance() Bitcoin {
	return wallet.balance
}

func (wallet *Wallet) Withdraw(bitcoin Bitcoin) error {
	if bitcoin > wallet.balance {
		return InsufficientFundsError
	}
	wallet.balance -= bitcoin
	return nil
}
