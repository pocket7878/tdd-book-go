package money

import (
	"testing"
)

func TestMultiplication(t *testing.T) {
	five := NewDollar(5)
	five.Times(2)
	if five.Amount() != 10 {
		t.Fatalf("Expected 10, got %d", five.Amount())
	}
}
