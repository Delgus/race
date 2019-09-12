package main

import (
	"fmt"
	//bank "github.com/delgus/race/bank/v1"
	//bank "github.com/delgus/race/bank/v2"
	//bank "github.com/delgus/race/bank/v3"
	//bank "github.com/delgus/race/bank/v4"
	//bank "github.com/delgus/race/bank/v5"
	//bank "github.com/delgus/race/bank/v6"
	bank "github.com/delgus/race/bank/v7"
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
