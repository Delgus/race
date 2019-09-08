package v5

import (
	"sync"
	"testing"
	"time"
)

func TestRaceDeposit(t *testing.T) {
	var group sync.WaitGroup
	group.Add(10000)
	for i := 0; i < 10000; i++ {
		go func() {
			Deposit(100)
			group.Done()
		}()
	}
	group.Wait()
	if Balance() != 1000000 {
		t.Errorf("unexpected balance: value - %d expect 1000000", Balance())
	}
}

func TestRaceWithDraw(t *testing.T) {
	setBalance(1000000)
	var group sync.WaitGroup
	group.Add(10000)
	for i := 0; i < 10000; i++ {
		go func() {
			WithDraw(100)
			group.Done()
		}()
	}
	group.Wait()
	if Balance() != 0 {
		t.Errorf("unexpected balance: value - %d expect 1000000", Balance())
	}
}

func TestRaceWithDraw2(t *testing.T) {
	setBalance(1000000)
	var group sync.WaitGroup
	group.Add(20000)
	for i := 0; i < 10000; i++ {
		go func() {
			if WithDraw(2000000) {
				t.Error("unexpected true for withdraw")
			}
			group.Done()
		}()
		go func() {
			if !WithDraw(100) {
				t.Error("unexpected false for withdraw")
			}
			time.Sleep(time.Nanosecond)
			group.Done()
		}()
	}
	group.Wait()
	if Balance() != 0 {
		t.Errorf("unexpected balance: value - %d expect 1000000", Balance())
	}
}
