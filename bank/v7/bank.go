package v6

import (
	"sync/atomic"
)

var balance int64

func Deposit(amount int64) {
	atomic.AddInt64(&balance, amount)
}

func Balance() int64 {
	return balance
}

func WithDraw(amount int64) bool {
	Deposit(-amount)
	if balance < 0 {
		Deposit(amount)
		return false
	}
	return true
}

func setBalance(count int64) {
	balance = count
}
