package class6_1

import "fmt"

type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

type Stringer interface {
	String() string
}

//Deposit Deposit
func (wallet *Wallet) Deposit(amount Bitcoin) {
	wallet.balance += amount
}

//GetBalance GetBalance
func (wallet *Wallet) GetBalance() Bitcoin {
	return wallet.balance
}

func (wallet *Wallet) Withdraw(amount Bitcoin) {
	wallet.balance -= amount
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
