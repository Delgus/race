package bank

import "sync"

var (
	mu      sync.Mutex
	balance int
)

func Deposit(amount int) {
	mu.Lock()
	deposit(amount)
	mu.Unlock()
}

func deposit(amount int) {
	balance = balance + amount
}

func Balance() int {
	mu.Lock()
	b := balance
	mu.Unlock()
	return b
}

func WithDraw(amount int) bool {
	mu.Lock()
	defer func() {
		mu.Unlock()
	}()
	deposit(-amount)
	if balance < 0 {
		deposit(amount)
		return false
	}
	return true
}

func setBalance(count int) {
	balance = count
}
