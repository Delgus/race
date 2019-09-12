### Bank v 7.0.0 

Седьмой релиз банка. Ядерный банк)

Функции:
Balance - получить баланс
Deposit - положить на счет
WithDraw - снять деньги со счета

тип баланса изменен на int64

Бенчмарк
```
goos: linux
goarch: amd64
pkg: github.com/delgus/race/bank/v7
BenchmarkBalance-4      200000000                6.97 ns/op            0 B/op          0 allocs/op
BenchmarkWithDraw-4     200000000                9.19 ns/op            0 B/op          0 allocs/op
PASS
ok      github.com/delgus/race/bank/v7  4.874s
```

А так можно было? Это вообще законно