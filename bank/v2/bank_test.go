package v2

import (
	"testing"
)

func TestDeposit(t *testing.T) {
	Deposit(100)
	expect := 100
	got := Balance()
	if got != expect {
		t.Errorf(
			"unexpected balance: value - %d expect %d", got, expect)
	}
}

func TestRaceDeposit(t *testing.T) {
	go Deposit(10)
	Deposit(10)
}

func BenchmarkBalance(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Deposit(1000)
	}
}
