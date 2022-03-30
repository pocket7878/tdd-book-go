package money

type Money struct {
	amount int
	name   string
}

type AmountGetter interface {
	Amount() int
}

type NameGetter interface {
	Name() string
}

type moneyEqualable interface {
	AmountGetter
	NameGetter
}

func (m *Money) Amount() int {
	return m.amount
}

func (m *Money) Name() string {
	return m.name
}

func (m *Money) Equals(other moneyEqualable) bool {
	return m.Name() == other.Name() && m.amount == other.Amount()
}
