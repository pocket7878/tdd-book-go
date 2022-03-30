package money

func NewDollar(amount int) *Money {
	return &Money{amount, "Dollar"}
}
