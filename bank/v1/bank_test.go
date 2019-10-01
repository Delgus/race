package v1

import (
	"testing"
)

func TestDeposit(t *testing.T) {
	setBalance(0)
	Deposit(100)
	expect := 100
	got := Balance()
	if got != expect {
		t.Errorf(
			"unexpected balance: value - %d expect %d", got, expect)
	}
}

func BenchmarkBalance(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Deposit(1000)
	}
}
