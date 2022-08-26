package pae

import (
	"testing"
)

func TestWallet(t *testing.T) {
	assertBalanceEqual := func(t *testing.T, wallet Wallet, want Bitcoin) {
		t.Helper()
		got := wallet.Balance()
		if got != want {
			t.Errorf("for %#v, got %s, want %s", wallet, got, want)
		}
	}
	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10.0))
		assertBalanceEqual(t, wallet, Bitcoin(10.0))
	})
	t.Run("withdrawl", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10.0))
		wallet.Withdrawl(Bitcoin(5.0))
		assertBalanceEqual(t, wallet, Bitcoin(5.0))
	})
	t.Run("withdrawl returns the amount of bitcoin", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10.0))
		want := wallet.Withdrawl(Bitcoin(5.0))
		assertBalanceEqual(t, wallet, want)
	})
	t.Run("withdrawl returns error if user is withdrawing too much", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(4.0))
		err := wallet.Withdrawl(Bitcoin(5.0))
		assertBalanceEqual(t, wallet, want)
		if err != nil {
			t.Errorf("expected to get an error")
		}
	})
}
