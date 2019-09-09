package wallet

import (
	"testing"
)

func TestWallet(t *testing.T) {
	assertBalance := func(t *testing.T, w Wallet, want Bitcoin) {
		t.Helper()
		got := w.Balance()
		if got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	}

	t.Run("Save", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(Bitcoin(10))

		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(10)}
		wallet.Withdraw(Bitcoin(5))

		assertBalance(t, wallet, Bitcoin(5))
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		stringBalance := Bitcoin(10)
		wallet := Wallet{stringBalance}
		err := wallet.Withdraw(Bitcoin(20))

		assertBalance(t, wallet, stringBalance)

		if err == nil {
			t.Errorf("wanted an error but didn't get one")
		}
	})
}
