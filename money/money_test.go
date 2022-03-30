package money

import (
	"testing"
)

func TestMultiplication(t *testing.T) {
	five := NewDollar(5)
	product := five.Times(2)
	if product.Amount() != 10 {
		t.Fatalf("Expected 10, got %d", product.Amount())
	}

	product = five.Times(3)
	if product.Amount() != 15 {
		t.Fatalf("Expected 15, got %d", product.Amount())
	}
}
