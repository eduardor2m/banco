package accounts

import "github.com/eduardor2m/banco/customers"

type SavingsAccount struct {
	Holder                                 customers.Customer
	AgencyNumber, AccountNumber, operation int
	balance                                float64
}

func (c *SavingsAccount) Withdraw(value float64) string {
	can := value > 0 && value <= c.balance

	if can {
		c.balance -= value
		return "Withdraw completed"
	} else {
		return "Insufficient balance"
	}

}

func (c *SavingsAccount) Deposit(value float64) (string, float64) {
	if value > 0 {
		c.balance += value
		return "Deposit completed", c.balance
	} else {
		return "Deposit value must be greater than zero", c.balance
	}
}

func (c *SavingsAccount) GetBalance() float64 {
	return c.balance
}
