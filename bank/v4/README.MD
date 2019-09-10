### Bank v 4.0.0

Четвертый релиз банка. Переписан полностью. Конвейер

Функции:
Balance - получить баланс
Deposit - положить на счет
WithDraw - снять деньги со счета

Бенчмарк
```
goos: linux
goarch: amd64
pkg: github.com/delgus/race/bank/v4
BenchmarkBalance-2      10000000               156 ns/op               0 B/op          0 allocs/op
BenchmarkWithDraw-2      5000000               277 ns/op               0 B/op          0 allocs/op
PASS
ok      github.com/delgus/race/bank/v4  3.405s
```
Поддерживать код стало не легче