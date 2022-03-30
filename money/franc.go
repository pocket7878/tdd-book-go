package money

type Franc struct {
	Money
}

func NewFranc(amount int) *Money {
	return newMoney(amount, "CHF")
}
