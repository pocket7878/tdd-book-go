package money

type Dollar struct{}

func NewDollar(amount int) *Dollar {
	return &Dollar{}
}

func (d *Dollar) Times(multiplier int) {
}

func (d *Dollar) Amount() int {
	return 0
}
