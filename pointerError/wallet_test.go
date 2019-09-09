package pointerError

import (
	"testing"
)

func TestWallet(t *testing.T) {
	wallet := Wallet{}
	wallet.Deposit(Bitcoin(10))
	got := wallet.Balance()
	want := Bitcoin(10)
	if got != want {
		t.Errorf("want %s, got %s", want, got)
	}
}
