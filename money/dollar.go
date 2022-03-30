package money

type Dollar struct {
	amount int
}

func NewDollar(amount int) *Dollar {
	return &Dollar{amount}
}

func (d *Dollar) Times(multiplier int) *Dollar {
	d.amount *= multiplier

	return nil
}

func (d *Dollar) Amount() int {
	return d.amount
}
