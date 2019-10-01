package v6

import (
	"sync"
	"testing"
)

func TestDeposit(t *testing.T) {
	setBalance(0)
	Deposit(100)
	var expect int64 = 100
	got := Balance()
	if got != expect {
		t.Errorf(
			"unexpected balance: value - %d expect %d", got, expect)
	}
}

func TestWithDrawSuccess(t *testing.T) {
	result := WithDraw(100)
	var expect int64 = 0
	got := Balance()
	if got != expect {
		t.Errorf("unexpected balance: value - %d expect %d", got, expect)
	}
	if !result {
		t.Error("unexpected result: value - false expect true")
	}
}

func TestWithDrawError(t *testing.T) {
	result := WithDraw(100)
	var expect int64 = 0
	got := Balance()
	if got != expect {
		t.Errorf("unexpected balance: value - %d expect %d", got, expect)
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
	var expect int64 = 100
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

func BenchmarkWithDraw(b *testing.B) {
	for i := 0; i < b.N; i++ {
		WithDraw(1000)
	}
}
