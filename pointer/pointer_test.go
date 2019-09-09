package pointer

import (
	"testing"
)



//TestWallet test wallet func
func TestWallet(t *testing.T){
	assertBalance := func(t *testing.T,wallet *Wallet,want Bitcoin){
		t.Helper()
		got := wallet.Balance()
		if got != want{
			t.Errorf("got %s want %s", got, want)
		}
	}

	t.Run("deposit",func(t *testing.T){
		wallet := &Wallet{}
		wallet.Deposit(Bitcoin(10))
		assertBalance(t,wallet,Bitcoin(10))
	})

	t.Run("withdraw",func(t *testing.T){
		wallet := &Wallet{balance: Bitcoin(20)}
		wallet.Withdraw(Bitcoin(10))
		assertBalance(t,wallet,Bitcoin(10))
	})

}