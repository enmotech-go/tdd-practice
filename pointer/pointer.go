package pointer

import "fmt"


type Bitcoin int

//Wallet define wallet struct
type Wallet struct {
	balance Bitcoin
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