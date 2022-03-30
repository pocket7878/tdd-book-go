package money

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMultiplication(t *testing.T) {
	five := NewDollar(5)
	assert.True(t, five.Times(2).Equals(NewDollar(10)))
	assert.True(t, five.Times(3).Equals(NewDollar(15)))
}

func TestEquality(t *testing.T) {
	assert.True(t, NewDollar(5).Equals(NewDollar(5)))
	assert.True(t, NewFranc(5).Equals(NewFranc(5)))

	assert.False(t, NewDollar(5).Equals(NewDollar(6)))
	assert.False(t, NewFranc(5).Equals(NewFranc(6)))

	assert.False(t, NewDollar(5).Equals(NewFranc(5)))
	assert.False(t, NewFranc(5).Equals(NewDollar(5)))
}

func TestFrancMultiplication(t *testing.T) {
	five := NewFranc(5)
	assert.True(t, five.Times(2).Equals(NewFranc(10)))
	assert.True(t, five.Times(3).Equals(NewFranc(15)))
}

func TestCurrency(t *testing.T) {
	assert.Equal(t, "USD", NewDollar(1).Currency())
	assert.Equal(t, "CHF", NewFranc(1).Currency())
}

func TestSimpleAddition(t *testing.T) {
	five := NewDollar(5)
	sum := five.Plus(five)
	bank := NewBank()
	reduced := bank.Reduce(sum, "USD")
	assert.True(t, NewDollar(10).Equals(reduced))
}

func TestPlusReturnsSum(t *testing.T) {
	five := NewDollar(5)
	result := five.Plus(five)
	sum := result.(*Sum)
	assert.True(t, five.Equals(sum.Augend()))
	assert.True(t, five.Equals(sum.Addend()))
}

func TestReduceSum(t *testing.T) {
	sum := NewSum(NewDollar(3), NewDollar(4))
	bank := NewBank()
	result := bank.Reduce(sum, "USD")
	assert.True(t, NewDollar(7).Equals(result))
}

func TestReduceMoney(t *testing.T) {
	bank := NewBank()
	result := bank.Reduce(NewDollar(1), "USD")
	assert.True(t, NewDollar(1).Equals(result))
}
