package v1

import (
	"testing"
)

func TestDeposit(t *testing.T) {
	setBalance(0)
	Deposit(100)
	if Balance() != 100 {
		t.Errorf(
			"unexpected balance: value - %d expect 100",
			Balance())
	}
}

func BenchmarkBalance(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Deposit(1000)
	}
}
