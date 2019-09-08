package v1

import (
	"testing"
)

func TestDeposit(t *testing.T) {
	Deposit(100)
	Balance()
	if Balance() != 100 {
		t.Errorf("unexpected balance: value - %d expect 100", Balance())
	}
}

func BenchmarkBalance(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Deposit(1000)
	}
}
