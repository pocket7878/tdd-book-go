package money

import "fmt"

type Bank struct {
	rates map[Pair]int
}

func NewBank() *Bank {
	return &Bank{
		rates: make(map[Pair]int),
	}
}

func (b *Bank) Reduce(source Expression, to string) *Money {
	return source.Reduce(b, to)
}

func (b *Bank) AddRate(from, to string, rate int) {
	b.rates[Pair{from, to}] = rate
}

func (b *Bank) Rate(from, to string) int {
	if from == to {
		return 1
	}
	rate, ok := b.rates[Pair{from, to}]
	if !ok {
		panic(fmt.Sprintf("No rate defined for %s -> %s", from, to))
	}

	return rate
}
