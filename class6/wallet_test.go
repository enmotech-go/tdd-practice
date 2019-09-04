package class6

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))

		fmt.Println("address of balance in test is", &wallet.balance)
		got := wallet.Balance()
		want := Bitcoin(10)
		if got != want {
			t.Errorf("got %d want %d ", got, want)
		}
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		wallet.Withdraw(Bitcoin(10))

		got := wallet.Balance()
		want := Bitcoin(10)
		if got != want {
			t.Errorf("got %d want %d ", got, want)
		}
	})

}
