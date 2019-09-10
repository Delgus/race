### Bank v 1.0.0

Первая версия банка.

Функции:
Balance - получить баланс
Deposit - положить на счет

Бенчмарк
```
goos: linux
goarch: amd64
pkg: github.com/delgus/race/bank/v1
BenchmarkBalance-2      300000000                4.24 ns/op            0 B/op          0 allocs/op
PASS
ok      github.com/delgus/race/bank/v1  1.707s
```

PS. Кажется безопасной ведь тесты проходят норм
```
alexey@alexey-K54L:~/go/src/github.com/delgus/race$ go test ./bank/v1 -race -cover
ok      github.com/delgus/race/bank/v1  1.014s  coverage: 100.0% of statements
```
Но на деле все не так. Вот этот тест например показал бы гонку

```go
func TestRaceDeposit(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func(){
			Deposit(10)
			wg.Done()
		}()
	}
	wg.Wait()
	if Balance() != 200 {
		t.Errorf("unexpected balance: value - %d expect 200", Balance())
	}
}
```