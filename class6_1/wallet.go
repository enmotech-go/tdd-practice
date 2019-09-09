package class6_1

import (
	"errors"
	"fmt"
)

type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

type Stringer interface {
	String() string
}

var InsufficientFundsError = errors.New("cannot withdraw, insufficient funds")

//Deposit Deposit
func (wallet *Wallet) Deposit(amount Bitcoin) {
	wallet.balance += amount
}

//GetBalance GetBalance
func (wallet *Wallet) GetBalance() Bitcoin {
	return wallet.balance
}

//Withdraw Withdraw
func (wallet *Wallet) Withdraw(amount Bitcoin) error {
	if amount > wallet.balance {
		return InsufficientFundsError
	}
	wallet.balance -= amount
	return nil
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
