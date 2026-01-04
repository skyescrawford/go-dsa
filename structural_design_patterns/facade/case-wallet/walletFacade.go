package main

import "fmt"

type WalletFacade struct {
	account      *Account
	wallet       *Wallet
	securityCode *SecurityCode
	notification *Notification
	ledger       *Ledger
}

func (w *WalletFacade) addMoney(id string, code int, amount int) error {
	fmt.Println("starting add money to wallet")

	err := w.account.checkAccount(id)
	if err != nil {
		return err
	}

	err = w.securityCode.checkCode(code)
	if err != nil {
		return err
	}

	w.wallet.creditBalance(amount)

	w.notification.sendWalletCreditNotification()

	w.ledger.makeEntry(id, "credit", amount)

	return nil
}

func (w *WalletFacade) deductMoney(id string, code int, amount int) error {
	fmt.Println("starting debit money from wallet")

	err := w.account.checkAccount(id)
	if err != nil {
		return err
	}

	err = w.securityCode.checkCode(code)
	if err != nil {
		return err
	}

	err = w.wallet.debitBalance(amount)
	if err != nil {
		return err
	}

	w.notification.sendWalletDebitNotification()

	w.ledger.makeEntry(id, "debit", amount)

	return nil
}

func (w *WalletFacade) checkBalance(id string, code int) {
	fmt.Println("starting check balance from wallet")

	w.account.checkAccount(id)

	w.securityCode.checkCode(code)

	fmt.Printf("your balance is %d\n", w.wallet.balance)

}

func newWalletFacade(id string, code int) *WalletFacade {
	return &WalletFacade{
		account:      newAccount(id),
		wallet:       newWallet(),
		securityCode: newSecurityCode(code),
		notification: &Notification{},
		ledger:       &Ledger{},
	}
}
