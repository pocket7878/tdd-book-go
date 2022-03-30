package money

type Money struct {
	amount int
}

type AmountGetter interface {
	Amount() int
}

func (m *Money) Amount() int {
	return m.amount
}

func (m *Money) Equals(other AmountGetter) bool {
	return m.amount == other.Amount()
}
