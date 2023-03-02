package main

import (
	"fmt"

	"github.com/eduardor2m/banco/accounts"
)

func payBill(account verifyAccount, value float64) {
	account.Withdraw(value)
}

type verifyAccount interface {
	Withdraw(value float64) string
}

func main() {
	accountDudu := accounts.SavingsAccount{}
	accountDudu.Deposit(100)
	payBill(&accountDudu, 60)

	fmt.Println(accountDudu.GetBalance())

	accountRicardo := accounts.CurrentAccount{}
	accountRicardo.Deposit(700)
	payBill(&accountRicardo, -360)

	fmt.Println(accountRicardo.GetBalance())

}
