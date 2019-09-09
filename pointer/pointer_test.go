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

	assertError := func(t *testing.T,err error,want string){
		t.Helper()
		if err == nil{
			t.Fatal("didn't get an error but wanted one")
		}
		if err.Error() != want {
			t.Errorf("got '%s', want '%s'", err, want)
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

	t.Run("Withdraw insufficient funds",func(t *testing.T){
		wallet := &Wallet{balance: Bitcoin(10)}
		err := wallet.Withdraw(Bitcoin(20))
		assertError(t,err,"cannot withdraw, insufficient funds")
	})
}