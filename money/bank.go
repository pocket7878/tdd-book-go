package money

import "fmt"

type Bank struct {
}

func NewBank() *Bank {
	return &Bank{}
}

func (b *Bank) Reduce(source Expression, to string) *Money {
	switch s := source.(type) {
	case *Sum:
		return s.Reduce(to)
	case *Money:
		return s
	default:
		panic(fmt.Sprintf("Unsupported source type: %T", s))
	}
}
