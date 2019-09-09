package pointerError

import (
	"testing"
)

func TestWallet(t *testing.T) {
	wallet := Wallet{}
	wallet.Deposit(10)
	got := wallet.Balance()
	want := 10
	if got != want {
		t.Errorf("want %d, got %d", want, got)
	}
}
