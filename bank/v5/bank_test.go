package v5

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

func TestWithDrawSuccess(t *testing.T) {
	setBalance(100)
	result := WithDraw(100)
	if Balance() != 0 {
		t.Errorf("unexpected balance: value - %d expect 100", Balance())
	}
	if !result {
		t.Error("unexpected result: value - false expect true")
	}
}

func TestWithDrawError(t *testing.T) {
	setBalance(0)
	result := WithDraw(100)
	if Balance() != 0 {
		t.Errorf("unexpected balance: value - %d expect 100", Balance())
	}
	if result {
		t.Error("unexpected result: value - true, expect false")
	}
}

func TestRaceDeposit(t *testing.T) {
	setBalance(0)
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			Deposit(10)
			wg.Done()
		}()
	}
	wg.Wait()
	if Balance() != 100 {
		t.Errorf("unexpected balance: value - %d expect 200", Balance())
	}
}

func BenchmarkBalance(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Deposit(1000)
	}
}

func BenchmarkWithDraw(b *testing.B) {
	for i := 0; i < b.N; i++ {
		WithDraw(1000)
	}
}
