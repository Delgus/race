### Bank v 6.0.0 

Шестой релиз банка. Переписан полностью на mutex

Функции:  
Balance - получить баланс  
Deposit - положить на счет  
WithDraw - снять деньги со счета  

Бенчмарк
```
goos: linux
goarch: amd64
pkg: github.com/delgus/race/bank/v6
BenchmarkBalance-2      30000000                41.3 ns/op             0 B/op          0 allocs/op
BenchmarkWithDraw-2     20000000               121 ns/op               0 B/op          0 allocs/op
PASS
ok      github.com/delgus/race/bank/v6  3.842s
```


Код легко поддерживать, скорость чумачечая