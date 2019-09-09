package class6_1

import "testing"

func TestWallet(t *testing.T) {

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		want := Bitcoin(10)
		assertBalance(t, &wallet, want)
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{Bitcoin(20)}
		_ = wallet.Withdraw(10)
		want := Bitcoin(10)
		assertBalance(t, &wallet, want)
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{Bitcoin(20)}
		_ = wallet.Withdraw(100)
		want := Bitcoin(10)
		got := wallet.GetBalance()
		if got != want {
			t.Log(got, want)
		}
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Bitcoin(100))

		assertBalance(t, &wallet, startingBalance)
		assertError(t, err, InsufficientFundsError)
	})
}

func assertError(t *testing.T, err error, want error) {
	if err == nil {
		t.Errorf("wanted an error but didnt get one")
	}
	if err != want {
		t.Errorf("got '%s', want '%s'", err.Error(), want)
	}
}

func assertBalance(t *testing.T, wallet *Wallet, want Bitcoin) {
	t.Helper()
	got := wallet.GetBalance()
	if got != want {
		t.Errorf("got %s want %s ", got, want)
	}
}
