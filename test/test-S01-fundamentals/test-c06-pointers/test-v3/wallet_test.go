package main

import (
	"testing"

	pointers "learn.go/S01-fundamentals/c06-pointers/v3"
)

func TestWallet(t *testing.T) {

	assertBalance := func(t testing.TB, wallet pointers.Wallet, want pointers.Bitcoin) {
		t.Helper()
		got := wallet.GetBalance()

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}

	assertError := func(t testing.TB, err error) {
		t.Helper()
		if err == nil {
			t.Error("wanted an error but didn't get one")
		}
	}

	t.Run("deposit", func(t *testing.T) {
		wallet := pointers.Wallet{}
		wallet.Deposit(pointers.Bitcoin(10))

		assertBalance(t, wallet, pointers.Bitcoin(10))
	})

	t.Run("withdraw with funds", func(t *testing.T) {
		wallet := pointers.Wallet{pointers.Bitcoin(20)}
		wallet.Withdraw(pointers.Bitcoin(10))

		assertBalance(t, wallet, pointers.Bitcoin(10))
	})

	t.Run("withdraw insufficient funds", func(t *testing.T) {
		startingBalance := pointers.Bitcoin(20)
		wallet := pointers.Wallet{startingBalance}
		err := wallet.Withdraw(pointers.Bitcoin(100))

		assertError(t, err)
		assertBalance(t, wallet, startingBalance)
	})
}
