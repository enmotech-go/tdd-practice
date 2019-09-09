package class6_1

import "testing"

func TestWallet(t *testing.T) {
	assertBalance := func(t *testing.T, wallet *Wallet, want Bitcoin) {
		t.Helper()
		got := wallet.GetBalance()
		if got != want {
			t.Errorf("got %s want %s ", got, want)
		}
	}
	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		want := Bitcoin(10)
		assertBalance(t, &wallet, want)
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{Bitcoin(20)}
		wallet.Withdraw(10)
		want := Bitcoin(10)
		assertBalance(t, &wallet, want)
	})
}
