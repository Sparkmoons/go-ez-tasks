## Рандомайзер

Функция генератора n случайных чисел в диапазоне от 0 до n-1. Но числа могут быть предсказуемы, поскольку генератор сделан на основе времени.

## Реализация

```go
func RandNumb(n int) <-chan int{
  r := rand.New(rand.NewSource(time.Now().UnixNano()))
  res := make(chan int)
  go func() {
    for j := 0; j < n; j++ {
      res <- r.Intn(n)
    }
    close(res)
  }()
  return res
}
```
