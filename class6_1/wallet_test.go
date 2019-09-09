package class6_1

import "testing"

func TestWallet(t *testing.T) {
	wallet := Wallet{}

	wallet.Deposit(10)

	got := wallet.GetBalance()
	want := 10
	if got != want {
		t.Errorf(" got %d want %d", got, want)
	}
}
