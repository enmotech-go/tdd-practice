package wallet

import "fmt"

type Bitcoin int

type Stringer interface {
	String() string
}

//define Wallet struct
type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) {
	//fmt.Println("address of balance", &w.balance)
	w.balance += amount
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (w *Wallet) Withdraw(amount Bitcoin) {
	w.balance -= amount
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
