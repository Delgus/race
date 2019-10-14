package v4

var (
	sema    = make(chan struct{}, 1) //семафор для блокировок
	balance int
)

//Deposit ...
func Deposit(amount int) {
	sema <- struct{}{}
	deposit(amount)
	<-sema
}

func deposit(amount int) {
	balance = balance + amount
}

//Balance ...
func Balance() int {
	sema <- struct{}{}
	b := balance
	<-sema
	return b
}

//WithDraw ...
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
