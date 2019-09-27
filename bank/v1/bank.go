package v1

var balance int

func Deposit(amount int) {
	balance = balance + amount
}

func Balance() int {
	return balance
}

func setBalance(amount int) {
	balance = amount
}
