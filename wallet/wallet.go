package wallet

//define Wallet struct
type Wallet struct {
	balance int
}

func (w *Wallet) Deposit(amount int) {
	//fmt.Println("address of balance", &w.balance)
	w.balance += amount
}

func (w *Wallet) Balance() int {
	return w.balance
}
