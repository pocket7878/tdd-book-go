package money

type Franc struct {
	amount int
}

func NewFranc(amount int) *Franc {
	return &Franc{amount}
}

func (d *Franc) Times(multiplier int) *Franc {
	return NewFranc(d.amount * multiplier)
}

func (d *Franc) Amount() int {
	return d.amount
}
