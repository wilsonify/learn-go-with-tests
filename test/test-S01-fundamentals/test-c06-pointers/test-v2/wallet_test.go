package main

import (
	"testing"

	pointers "learn.go/S01-fundamentals/c06-pointers/v2"
)

func TestWallet(t *testing.T) {

	assertBalance := func(t testing.TB, wallet pointers.Wallet, want pointers.Bitcoin) {
		t.Helper()
		got := wallet.GetBalance()

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}

	t.Run("deposit", func(t *testing.T) {
		wallet := pointers.Wallet{}
		wallet.Deposit(pointers.Bitcoin(10))
		assertBalance(t, wallet, pointers.Bitcoin(10))
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := pointers.Wallet{Balance: pointers.Bitcoin(20)}
		wallet.Withdraw(10)
		assertBalance(t, wallet, pointers.Bitcoin(10))
	})

}
