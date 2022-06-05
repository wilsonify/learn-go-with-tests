package pointers

import "fmt"

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

// Balance returns the number of Bitcoin a wallet has.
func (w *Wallet) GetBalance() Bitcoin {
	return w.Balance
}
