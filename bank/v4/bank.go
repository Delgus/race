package v4

type money struct {
	count int
}

var balances = make(chan *money, 1)

func Deposit(amount int) {
	balance := <-balances
	balance.count += amount
	balances <- balance
}

func Balance() int {
	balance := <-balances
	balances <- balance
	return balance.count
}

func WithDraw(amount int) bool {
	balance := <-balances
	defer func() {
		balances <- balance
	}()
	balance.count -= amount
	if balance.count < 0 {
		balance.count += amount
		return false
	}
	return true
}

func init() {
	balances <- new(money)
}

func setBalance(amount int) {
	<-balances
	balances <- &money{count: amount}
}
