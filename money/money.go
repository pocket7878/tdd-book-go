package money

import "fmt"

type Money struct {
	amount   int
	currency string
}

func newMoney(amount int, currency string) *Money {
	return &Money{amount, currency}
}

type AmountGetter interface {
	Amount() int
}

type CurrencyGetter interface {
	Currency() string
}

type moneyEqualable interface {
	AmountGetter
	CurrencyGetter
}

func (m *Money) Amount() int {
	return m.amount
}

func (m *Money) Currency() string {
	return m.currency
}

func (m *Money) Times(multiplier int) *Money {
	return &Money{m.amount * multiplier, m.currency}
}

func (m *Money) Plus(added *Money) Expression {
	return NewSum(m, added)
}

func (m *Money) Equals(other moneyEqualable) bool {
	return m.Currency() == other.Currency() && m.Amount() == other.Amount()
}

func (m *Money) String() string {
	return fmt.Sprintf("%d %s", m.Amount(), m.Currency())
}

func (m *Money) Reduce(bank *Bank, to string) *Money {
	rate := bank.Rate(m.Currency(), to)
	return newMoney(m.Amount()/rate, to)
}
