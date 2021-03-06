package v3

var deposits = make(chan int)         // Отправление вклада
var balances = make(chan int)         // Получение баланса
var withdraws = make(chan int)        //Снятие со счета
var withdrawSuccess = make(chan bool) //успешное или нет снятие со счета

//Deposit ...
func Deposit(amount int) {
	deposits <- amount
}

//Balance ...
func Balance() int {
	return <-balances
}

//WithDraw ...
func WithDraw(amount int) bool {
	withdraws <- amount
	return <-withdrawSuccess
}

func teller() {
	var balance int // balance ограничен go-подпрограммой teller
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
		case amount := <-withdraws:
			balance -= amount
			if balance < 0 {
				balance += amount
				withdrawSuccess <- false
			} else {
				withdrawSuccess <- true
			}
		}
	}
}

func init() {
	go teller() // Запуск управляющей go-подпрограммы
}
