package pointer

import "testing"

func TestWallet(t *testing.T) {
	wallet := Wallet{}
	wallet.Deposit(10)

	got := wallet.Balance()
	want := Bitcoin(10)
	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
