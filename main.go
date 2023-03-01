package main

import (
	"fmt"

	"github.com/eduardor2m/banco/accounts"
	"github.com/eduardor2m/banco/customers"
)

func main() {
	accountExemple := accounts.CurrentAccount{}

	accountExemple.Holder = customers.Customer{
		Name:       "Eduardo",
		CPF:        "123.456.789-10",
		Profession: "Developer",
	}

	accountExemple.Deposit(100)
	accountExemple.Withdraw(50)

	fmt.Println(accountExemple.GetBalance())

}
