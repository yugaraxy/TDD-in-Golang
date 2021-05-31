package money

import (
	"testing"
)

func TestMoneyMultiplication(t *testing.T) {
	five := Dollar(5)
	if Dollar(10) != five.Times(2) {
		t.Errorf("Dollar(5).times(2) should be Dollar(10)")
	}

	if Dollar(15) != five.Times(3) {
		t.Errorf("Dollar(5).times(3) should be Dollar(15)")
	}
}

func TestMoneyCurrency(t *testing.T) {
	if "USD" != Dollar(1).currency() {
		t.Errorf("Dollar's currency should be USD")
	}

	if "CHF" != Franc(1).currency() {
		t.Errorf("Franc's currency should be USD")
	}
}

func TestMoneyEquality(t *testing.T) {
	if !Dollar(5).Equals(Dollar(5)) {
		t.Errorf("dollar equivalence is lost")
	}

	if Dollar(5).Equals(Dollar(6)) {
		t.Errorf("dollar equivalence is lost")
	}

	if Dollar(5).Equals(Franc(5)) {
		t.Errorf("dollar(5) doesn't equal to franc(5)")
	}
}
