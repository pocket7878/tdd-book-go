package money

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

func (m *Money) Equals(other moneyEqualable) bool {
	return m.Currency() == other.Currency() && m.Amount() == other.Amount()
}
