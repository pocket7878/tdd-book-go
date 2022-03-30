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

func (s *Sum) Reduce(bank *Bank, to string) *Money {
	amount := s.augend.Amount() + s.addend.Amount()
	return newMoney(amount, to)
}
