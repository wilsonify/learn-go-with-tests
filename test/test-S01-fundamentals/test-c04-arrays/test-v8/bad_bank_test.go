package main

import (
	"testing"

	arrays "learn.go/S01-fundamentals/c04-arrays/v8"
)

func TestBadBank(t *testing.T) {
	var (
		riya  = arrays.Account{Name: "Riya", Balance: 100}
		chris = arrays.Account{Name: "Chris", Balance: 75}
		adil  = arrays.Account{Name: "Adil", Balance: 200}

		transactions = []arrays.Transaction{
			arrays.NewTransaction(chris, riya, 100),
			arrays.NewTransaction(adil, chris, 25),
		}
	)

	newBalanceFor := func(account arrays.Account) float64 {
		return arrays.NewBalanceFor(account, transactions).Balance
	}

	AssertEqual(t, newBalanceFor(riya), 200)
	AssertEqual(t, newBalanceFor(chris), 0)
	AssertEqual(t, newBalanceFor(adil), 175)
}
