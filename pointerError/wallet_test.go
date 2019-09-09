package pointerError

import (
	"testing"
)

func TestWallet(t *testing.T) {
	assert := func(t *testing.T, got, want Bitcoin) {
		t.Helper()
		if got != want {
			t.Errorf("got %s, but want %s", got, want)
		}
	}
	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		got := wallet.Balance()
		want := Bitcoin(10)
		assert(t, got, want)
	})

	t.Run("withdraw", func(t *testing.T) {
		w := Wallet{Bitcoin(20)}
		w.Withdraw(10)
		got := w.Balance()
		want := Bitcoin(10)
		assert(t, got, want)
	})
	assertErr := func(t *testing.T, err error, want string) {
		if err == nil {
			t.Fatalf("want an error but do not get one")
		}
		if err.Error() != want {
			t.Errorf("got %s, want %s", err, want)
		}
	}
	t.Run("withdraw insufficient funds", func(t *testing.T) {
		w := Wallet{Bitcoin(20)}
		err := w.Withdraw(100)
		assert(t, w.Balance(), Bitcoin(20))
		assertErr(t, err, "cannot withdraw, insufficient funds")
	})
}
