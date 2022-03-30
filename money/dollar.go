package money

func NewDollar(amount int) *Money {
	return newMoney(amount, "USD")
}
