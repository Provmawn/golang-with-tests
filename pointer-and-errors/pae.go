package pae

import (
	"errors"
	"fmt"
)

var ErrInsufficientFunds = errors.New("cannot withdrawl, insufficient funds.")

type Bitcoin float64

func (b Bitcoin) String() string {
	return fmt.Sprintf("%g BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (w *Wallet) Withdrawl(amount Bitcoin) error {
	if amount > w.Balance() {
		return ErrInsufficientFunds
	}
	w.balance -= amount
	return nil
}
