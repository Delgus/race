package main

import (
	"fmt"
	//bank "github.com/delgus/race/bank/v1"
	bank "github.com/delgus/race/bank/v2"
	"sync"
)

func main() {
	var group sync.WaitGroup
	group.Add(10000)
	for i := 0; i < 10000; i++ {
		go func() {
			bank.Deposit(100)
			group.Done()
		}()
	}
	group.Wait()
	fmt.Println(bank.Balance())
}
