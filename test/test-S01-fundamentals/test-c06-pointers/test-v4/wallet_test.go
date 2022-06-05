package main

import (
	"testing"

	pointers "learn.go/S01-fundamentals/c06-pointers/v4"
)

func TestWallet(t *testing.T) {

	t.Run("deposit", func(t *testing.T) {
		wallet := pointers.Wallet{}
		wallet.Deposit(pointers.Bitcoin(10))

		assertBalance(t, wallet, pointers.Bitcoin(10))
	})

	t.Run("withdraw with funds", func(t *testing.T) {
		wallet := pointers.Wallet{pointers.Bitcoin(20)}
		err := wallet.Withdraw(pointers.Bitcoin(10))

		assertBalance(t, wallet, pointers.Bitcoin(10))
		assertNoError(t, err)
	})

	t.Run("withdraw insufficient funds", func(t *testing.T) {
		startingBalance := pointers.Bitcoin(20)
		wallet := pointers.Wallet{startingBalance}
		err := wallet.Withdraw(pointers.Bitcoin(100))

		assertBalance(t, wallet, startingBalance)
		assertError(t, err, pointers.ErrInsufficientFunds)
	})
}

func assertBalance(t testing.TB, wallet pointers.Wallet, want pointers.Bitcoin) {
	t.Helper()
	got := wallet.GetBalance()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertNoError(t testing.TB, got error) {
	t.Helper()
	if got != nil {
		t.Fatal("got an error but didn't want one")
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("didn't get an error but wanted one")
	}

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
