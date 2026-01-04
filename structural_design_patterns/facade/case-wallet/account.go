package main

import "errors"

type Account struct {
	id string
}

func (a *Account) checkAccount(id string) error {
	if id != a.id {
		return errors.New("invalid id")
	}
	return nil
}

func newAccount(id string) *Account {
	return &Account{id: id}
}
