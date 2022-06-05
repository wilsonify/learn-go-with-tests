package pointers

import (
	"errors"
	"fmt"
)

// Bitcoin represents a number of Bitcoins.
type Bitcoin int

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

// Wallet stores the number of Bitcoin someone owns.
type Wallet struct {
	Balance Bitcoin
}

// Deposit will add some Bitcoin to a wallet.
func (w *Wallet) Deposit(amount Bitcoin) {
	w.Balance += amount
}

// ErrInsufficientFunds means a wallet does not have enough Bitcoin to perform a withdraw.
var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

// Withdraw subtracts some Bitcoin from the wallet, returning an error if it cannot be performed.
func (w *Wallet) Withdraw(amount Bitcoin) error {

	if amount > w.Balance {
		return ErrInsufficientFunds
	}

	w.Balance -= amount
	return nil
}

// Balance returns the number of Bitcoin a wallet has.
func (w *Wallet) GetBalance() Bitcoin {
	return w.Balance
}
