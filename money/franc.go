package money

type Franc struct {
	Money
}

func NewFranc(amount int) *Franc {
	return &Franc{Money{amount}}
}

func (d *Franc) Times(multiplier int) *Franc {
	return NewFranc(d.amount * multiplier)
}
