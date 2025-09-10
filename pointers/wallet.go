package main

import (
	"errors"
	"fmt"
)

var InsufficientFundError = errors.New("cannot withdraw, insufficientFund")

type Bitcom int

type Stringer interface {
	String() string
}

func (b Bitcom) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance Bitcom
}

func (w *Wallet) Deposit(amount Bitcom) {
	w.balance += amount
}

func (w *Wallet) Balance() Bitcom {
	return w.balance
}

func (w *Wallet) Withdraw(amount Bitcom) error {
	if amount > w.balance {
		return InsufficientFundError
	}
	w.balance -= amount
	return nil
}
