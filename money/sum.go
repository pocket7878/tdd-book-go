package money

type Sum struct {
	augend *Money
	addend *Money
}

func NewSum(augend *Money, addend *Money) *Sum {
	return &Sum{augend, addend}
}

func (s *Sum) Augend() *Money {
	return s.augend
}

func (s *Sum) Addend() *Money {
	return s.addend
}
