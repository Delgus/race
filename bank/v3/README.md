### Bank v 3.0.0

Третий релиз банка. Добавлена возможность снимать деньги со счета

Функции:  
Balance - получить баланс  
Deposit - положить на счет  
WithDraw - снять деньги со счета  

Бенчмарк  
```
goos: linux
goarch: amd64
pkg: github.com/delgus/race/bank/v3
BenchmarkBalance-2       2000000               985 ns/op               0 B/op          0 allocs/op
BenchmarkWithDraw-2      1000000              1701 ns/op               0 B/op          0 allocs/op
PASS
ok      github.com/delgus/race/bank/v3  4.685s
```
Поддерживать код стало значительно тяжелее, а ведь у нас всего 3 функции