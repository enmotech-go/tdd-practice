package pointer

import (
	"fmt"
	"testing"
)



//TestWallet test wallet func
func TestWallet(t *testing.T){
	wallet := &Wallet{}
	wallet.Deposit(10)
	got := wallet.Balance()
	want := Bitcoin(10)
	fmt.Println("address of balance in test is", &wallet.balance)
	if got != want{
		t.Errorf("got %d want %d",got ,want)
	}
}