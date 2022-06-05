package main

import (
	"testing"

	pointers "learn.go/S01-fundamentals/c06-pointers/v1"
)

func TestWallet(t *testing.T) {

	wallet := pointers.Wallet{}

	wallet.Deposit(pointers.Bitcoin(10))

	got := wallet.GetBalance()

	want := pointers.Bitcoin(10)

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}
