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
	assertError := func(t *testing.T, got, want error) {
		t.Helper()
		if got == nil {
			t.Error("wanted an error but didn't get one")
		}

		if got != want {
			t.Errorf(`got "%s", want "%s"`, got, want)
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
		assertError(t, err, InsufficientFunds)
	})

	t.Run("test string Bitcoin", func(t *testing.T) {
		bitcoin := Bitcoin(10)
		want := "10 BTC"
		got := bitcoin.String()
		if want != got {
			t.Errorf("got %s, want %s", got, want)
		}
	})
}
