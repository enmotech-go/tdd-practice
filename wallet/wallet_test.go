package wallet

import (
	"testing"
)

func TestWallet(t *testing.T) {
	wallet := Wallet{}

	wallet.Deposit(10)

	got := wallet.Balance()
	//fmt.Println("address of balance", &wallet.balance)
	want := 10

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
