package pointerError

import (
	"testing"
)

func TestWallet(t *testing.T) {
	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		got := wallet.Balance()
		want := Bitcoin(10)
		if got != want {
			t.Errorf("want %s, got %s", want, got)
		}
	})

	t.Run("withdraw", func(t *testing.T) {
		w := Wallet{Bitcoin(20)}
		w.Withdraw(10)
		got := w.Balance()
		want := Bitcoin(10)
		if got != want {
			t.Errorf("want %s, got %s", want, got)
		}
	})
}
