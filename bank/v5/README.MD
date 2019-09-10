### Bank v 5.0.0

Пятый релиз банка. Переписан полностью. Семафор

Функции:
Balance - получить баланс
Deposit - положить на счет
WithDraw - снять деньги со счета

Бенчмарк
```
goos: linux
goarch: amd64
pkg: github.com/delgus/race/bank/v5
BenchmarkBalance-2      10000000               128 ns/op               0 B/op          0 allocs/op
BenchmarkWithDraw-2     10000000               216 ns/op               0 B/op          0 allocs/op
PASS
ok      github.com/delgus/race/bank/v5  3.841s
```
Поддерживать код стало легче. Если ты конечно понимаешь что такое каналы