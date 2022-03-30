package money

type Bank struct {
}

func NewBank() *Bank {
	return &Bank{}
}

func (b *Bank) Reduce(source Expression, to string) *Money {
	sum, ok := source.(*Sum)
	if !ok {
		panic("Expression is not a Sum")
	}
	return sum.Reduce(to)
}
