package v5

var (
	sema    = make(chan struct{}, 1) //семафор для блокировок
	balance int
)

func Deposit(amount int) {
	sema <- struct{}{}
	deposit(amount)
	<-sema
}

func deposit(amount int) {
	balance = balance + amount
}

func Balance() int {
	sema <- struct{}{}
	b := balance
	<-sema
	return b
}

func WithDraw(amount int) bool {
	sema <- struct{}{}
	defer func() {
		<-sema
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
