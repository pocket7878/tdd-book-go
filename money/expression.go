package money

type Expression interface {
	Reduce(bank *Bank, to string) *Money
	Times(multiplier int) Expression
}
