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
BenchmarkBalance-2      100000000               15.4 ns/op             0 B/op          0 allocs/op
BenchmarkWithDraw-2     50000000                25.5 ns/op             0 B/op          0 allocs/op
PASS
ok      github.com/delgus/race/bank/v7  2.876s
```

А так можно было? Это вообще законно