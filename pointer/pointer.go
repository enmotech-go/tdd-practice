package pointer

import "fmt"

//Wallet define wallet struct
type Wallet struct {
	balance int
}


//Deposit deposit bitcoin to wallet
func (w *Wallet) Deposit(amount int){
	w.balance += amount
}

//Balance get bitcoin from wallet
func (w *Wallet) Balance() (balance int){
 	balance = w.balance
	fmt.Println("address of balance in Deposit is", &w.balance)
	return
}