package pointer

import "fmt"


type Bitcoin int

//Wallet define wallet struct
type Wallet struct {
	balance Bitcoin
}

type Stringer interface {
	String() string
}


//Deposit deposit bitcoin to wallet
func (w *Wallet) Deposit(amount Bitcoin){
	w.balance += amount
}

//Balance get bitcoin from wallet
func (w *Wallet) Balance() (balance Bitcoin){
 	balance = w.balance
	fmt.Println("address of balance in Deposit is", &w.balance)
	return
}

//Withdraw get bitcoin from wallet
func(w *Wallet) Withdraw(amount Bitcoin){
	w.balance -= amount
}

func (b Bitcoin) String() (str string){
	str = fmt.Sprintf("%d BTC",b)
	return
}
