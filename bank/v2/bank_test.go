package v2

import (
	"sync"
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
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			Deposit(10)
			wg.Done()
		}()
	}
	wg.Wait()
	expect := 200 //из прежнего кейса на балансе осталось 100
	got := Balance()
	if got != expect {
		t.Errorf("unexpected balance: value - %d expect %d", got, expect)
	}
}

func BenchmarkBalance(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Deposit(1000)
	}
}
