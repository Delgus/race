package v2

import (
	"sync"
	"testing"
)

func TestDeposit(t *testing.T) {
	Deposit(100)
	if Balance() != 100 {
		t.Errorf("unexpected balance: value - %d expect 100", Balance())
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
	if Balance() != 200 {
		t.Errorf("unexpected balance: value - %d expect 200", Balance())
	}
}

func BenchmarkBalance(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Deposit(1000)
	}
}
