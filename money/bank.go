package money

type Bank struct {
}

func NewBank() *Bank {
	return &Bank{}
}

func (b *Bank) Reduce(source Expression, to string) *Money {
	sum, ok := source.(*Sum)
	if ok {
		return sum.Reduce(to)
	}
	money, ok := source.(*Money)
	if ok {
		return money
	}
	panic("Unsupported source type")
}
