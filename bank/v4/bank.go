package v4

type purse struct {
	money int
}

var balances = make(chan *purse, 1)

//Deposit ...
func Deposit(amount int) {
	balance := <-balances
	balance.money += amount
	balances <- balance
}

//Balance ...
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
	balances <- new(purse)
}

func setBalance(amount int) {
	<-balances
	balances <- &purse{money: amount}
}
