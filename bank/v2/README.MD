### Bank v 2.0.0

Безопасная версия банка.
Пример банка, переписанный с переменной balance , которая ограничивается монитором teller :

Функции:
Balance - получить баланс
Deposit - положить на счет

Бенчмарк
```
goos: linux
goarch: amd64
pkg: github.com/delgus/race/bank/v2
BenchmarkBalance-2       2000000               982 ns/op               0 B/op          0 allocs/op
PASS
ok      github.com/delgus/race/bank/v2  2.929s
```