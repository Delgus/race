package v5

import "sync"

var (
	mu      sync.Mutex
	balance int
)

//Deposit ...
func Deposit(amount int) {
	mu.Lock()
	deposit(amount)
	mu.Unlock()
}

//Balance ...
func Balance() int {
	mu.Lock()
	b := balance
	mu.Unlock()
	return b
}

//WithDraw ...
func WithDraw(amount int) bool {
	mu.Lock()
	defer mu.Unlock()
	withDraw(amount)
	if balance < 0 {
		deposit(amount)
		return false
	}
	return true
}

func deposit(amount int) {
	balance = balance + amount
}

func withDraw(amount int) {
	balance = balance - amount
}

func setBalance(count int) {
	balance = count
}
