package v4

type Money struct {
	money int
}

var balances = make(chan *Money, 1)

func Deposit(amount int) {
	balance := <-balances
	balance.money += amount
	balances <- balance
}

func Balance() int {
	balance := <-balances
	balances <- balance
	return balance.money
}

func WithDraw(amount int) bool {
	balance := <-balances
	defer func() {
		balances <- balance
	}()
	balance.money -= amount
	if balance.money < 0 {
		balance.money += amount
		return false
	}
	return true
}

func init() {
	balances <- new(Money)
}
