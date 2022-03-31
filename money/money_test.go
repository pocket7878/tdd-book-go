package money

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMultiplication(t *testing.T) {
	five := NewDollar(5)
	assert.Equal(t, five.Times(2), NewDollar(10))
	assert.Equal(t, five.Times(3), NewDollar(15))
}

func TestEquality(t *testing.T) {
	assert.Equal(t, NewDollar(5), NewDollar(5))
	assert.Equal(t, NewFranc(5), NewFranc(5))

	assert.NotEqual(t, NewDollar(5), NewDollar(6))
	assert.NotEqual(t, NewFranc(5), NewFranc(6))

	assert.NotEqual(t, NewDollar(5), NewFranc(5))
	assert.NotEqual(t, NewFranc(5), NewDollar(5))
}

func TestFrancMultiplication(t *testing.T) {
	five := NewFranc(5)
	assert.Equal(t, five.Times(2), NewFranc(10))
	assert.Equal(t, five.Times(3), NewFranc(15))
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
	assert.Equal(t, NewDollar(10), reduced)
}

func TestPlusReturnsSum(t *testing.T) {
	five := NewDollar(5)
	result := five.Plus(five)
	sum := result.(*Sum)
	assert.Equal(t, five, sum.Augend())
	assert.Equal(t, five, sum.Addend())
}

func TestReduceSum(t *testing.T) {
	sum := NewSum(NewDollar(3), NewDollar(4))
	bank := NewBank()
	result := bank.Reduce(sum, "USD")
	assert.Equal(t, NewDollar(7), result)
}

func TestReduceMoney(t *testing.T) {
	bank := NewBank()
	result := bank.Reduce(NewDollar(1), "USD")
	assert.Equal(t, NewDollar(1), result)
}

func TestReduceMoneyDifferentCurrency(t *testing.T) {
	bank := NewBank()
	bank.AddRate("CHF", "USD", 2)
	result := bank.Reduce(NewFranc(2), "USD")
	assert.Equal(t, NewDollar(1), result)
}

func TestIdentityRate(t *testing.T) {
	assert.Equal(t, 1, NewBank().Rate("USD", "USD"))
}

func TestMixedAddition(t *testing.T) {
	fiveBucks := NewDollar(5)
	tenFrancs := NewFranc(10)
	bank := NewBank()
	bank.AddRate("CHF", "USD", 2)
	result := bank.Reduce(fiveBucks.Plus(tenFrancs), "USD")
	assert.Equal(t, NewDollar(10), result)
}
