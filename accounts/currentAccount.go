package accounts

import "github.com/eduardor2m/banco/customers"

type CurrentAccount struct {
	Holder        customers.Customer
	AgencyNumber  int
	AccountNumber int
	balance       float64
}

func (c *CurrentAccount) Withdraw(value float64) string {
	can := value > 0 && value <= c.balance

	if can {
		c.balance -= value
		return "Withdraw completed"
	} else {
		return "Insufficient balance"
	}

}

func (c *CurrentAccount) Deposit(value float64) (string, float64) {
	if value > 0 {
		c.balance += value
		return "Deposit completed", c.balance
	} else {
		return "Deposit value must be greater than zero", c.balance
	}
}

func (c *CurrentAccount) Transfer(value float64, accountDestination *CurrentAccount) bool {
	if value > 0 && value <= c.balance {
		c.balance -= value
		accountDestination.Deposit(value)
		return true
	} else {
		return false
	}
}

func (c *CurrentAccount) GetBalance() float64 {
	return c.balance
}
