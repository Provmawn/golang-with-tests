package pae

import "fmt"

type Bitcoin float64

type Wallet struct {
	balance Bitcoin
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%g BTC", b)
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (w *Wallet) Withdrawl(amount Bitcoin) Bitcoin {
	if amount > w.balance {
		panic("withdrew past balance")
	}
	w.balance -= amount
	return amount
}
