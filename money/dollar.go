package money

type Dollar struct {
	amount int
}

func NewDollar(amount int) *Dollar {
	return &Dollar{10}
}

func (d *Dollar) Times(multiplier int) {
}

func (d *Dollar) Amount() int {
	return d.amount
}
