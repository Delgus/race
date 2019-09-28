package v1

var balance int

//Deposit ...
func Deposit(amount int) {
	balance = balance + amount
}

//Balance ...
func Balance() int {
	return balance
}

func setBalance(amount int) {
	balance = amount
}
