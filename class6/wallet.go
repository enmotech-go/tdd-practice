package class6

import (
	"errors"
)

type Bitcoin int
type Wallet struct {
	balance Bitcoin
}

func (wallet *Wallet) Deposit(amount Bitcoin) {
	wallet.balance += amount
}

func (wallet *Wallet) Balance() Bitcoin {
	return wallet.balance
}

func (wallet *Wallet) Withdraw(bitcoin Bitcoin) error {
	if bitcoin > wallet.balance {
		return errors.New("cannot withdraw, insufficient funds")
	}
	wallet.balance -= bitcoin
	return nil
}
