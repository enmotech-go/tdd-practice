package pointer

import "testing"

func TestWallet(t *testing.T) {
	assertBalance := func(t *testing.T, wallet Wallet, want Bitcoin) {
		t.Helper()
		got := wallet.Balance()
		if got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	}

	t.Run("test deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(10)
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("test withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		wallet.Withdraw(10)
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("test insufficient", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(10)}
		err := wallet.Withdraw(20)
		assertBalance(t, wallet, Bitcoin(10))

		if err == nil {
			t.Errorf("wanted an error but didn't get one")
		}
	})
}
