package money

import (
	"testing"
)

func TestMultiplication(t *testing.T) {
	five := NewDollar(5)
	product := five.Times(2)
	if !product.Equals(NewDollar(10)) {
		t.Fatalf("Expected 10, got %d", product.Amount())
	}

	product = five.Times(3)
	if !product.Equals(NewDollar(15)) {
		t.Fatalf("Expected 15, got %d", product.Amount())
	}
}

func TestEquality(t *testing.T) {
	five := NewDollar(5)
	fiveB := NewDollar(5)

	if !five.Equals(fiveB) {
		t.Fatalf("Expected true, got false")
	}
}

func TestFrancMultiplication(t *testing.T) {
	five := NewFranc(5)
	product := five.Times(2)
	if !product.Equals(NewFranc(10)) {
		t.Fatalf("Expected 10, got %d", product.Amount())
	}

	product = five.Times(3)
	if !product.Equals(NewFranc(15)) {
		t.Fatalf("Expected 15, got %d", product.Amount())
	}
}
