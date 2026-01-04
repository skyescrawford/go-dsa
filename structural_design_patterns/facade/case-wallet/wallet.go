package main

import "errors"

type Wallet struct {
	balance int
}

func (w *Wallet) creditBalance(amount int) {
	w.balance += amount
}

func (w *Wallet) debitBalance(amount int) error {
	if amount > w.balance {
		return errors.New("insufficient balance")
	}
	w.balance -= amount
	return nil
}

func newWallet() *Wallet {
	return &Wallet{}
}
