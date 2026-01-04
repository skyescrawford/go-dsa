package main

func main() {
	wallet := newWalletFacade("skyes07", 121212)

	wallet.addMoney("skyes07", 121212, 100000)

	wallet.deductMoney("skyes07", 121212, 50000)

	wallet.checkBalance("skyes07", 121212)
}
