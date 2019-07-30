package bank

import (
	"sync"
	"testing"
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
		t.Errorf("unexpected balance: value - %d expect 1000000", Balance() )
	}
}
