package money

type Dollar struct {
	amount int
}

func NewDollar(amount int) *Dollar {
	return &Dollar{amount}
}

func (d *Dollar) Times(multiplier int) {
	d.amount *= multiplier
}

func (d *Dollar) Amount() int {
	return d.amount
}
