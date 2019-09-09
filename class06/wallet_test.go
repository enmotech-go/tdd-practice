package class06

import (
	"testing"
)

func TestWallet(t *testing.T) {
	wallet := Wallet{}

	wallet.Deposit(Bitcoin(10))

	got := wallet.Balance()
	want := Bitcoin(10)

	if got.String() != want.String() {
		t.Errorf("Balance() = '%v', want '%v'", got.String(), want.String())
	}
}
