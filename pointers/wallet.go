package pointers

import (
	"errors"
	"fmt"
)

// Bitcoin currency
type Bitcoin int

// String the output format
func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

// Wallet struct
type Wallet struct {
	balance Bitcoin
}

// Deposit deposits money to Wallet
func (w *Wallet) Deposit(amount Bitcoin) {
	fmt.Printf("address of balance in Deposit is %v \n", &w.balance)
	w.balance += amount
}

// Balance get balance in Wallet
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

// ErrInsufficientFunds error message
var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

// Withdraw some moeny from wallet
func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrInsufficientFunds
	}

	w.balance -= amount
	return nil
}
