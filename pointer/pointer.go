package pointer

import (
	"fmt"

	"github.com/pkg/errors"
)

type Bitcoin int

type Stringer interface {
	String() string
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

//Wallet define wallet struct
type Wallet struct {
	balance Bitcoin
}

//Deposit deposit bitcoin to wallet
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

//Balance get bitcoin from wallet
func (w *Wallet) Balance() (balance Bitcoin) {
	balance = w.balance
	fmt.Println("address of balance in Deposit is", &w.balance)
	return
}

var InsufficientFunds = errors.New("cannot withdraw, insufficient funds")

//Withdraw get bitcoin from wallet
func (w *Wallet) Withdraw(amount Bitcoin) (err error) {
	if amount > w.balance {
		return InsufficientFunds
	}
	w.balance -= amount
	return
}
