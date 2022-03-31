package money

type Sum struct {
	augend Expression
	addend Expression
}

func NewSum(augend Expression, addend Expression) *Sum {
	return &Sum{augend, addend}
}

func (s *Sum) Augend() Expression {
	return s.augend
}

func (s *Sum) Addend() Expression {
	return s.addend
}

func (s *Sum) Reduce(bank *Bank, to string) *Money {
	amount := s.augend.Reduce(bank, to).Amount() + s.addend.Reduce(bank, to).Amount()
	return newMoney(amount, to)
}

func (s *Sum) Plus(added Expression) Expression {
	return nil
}
