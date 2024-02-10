## Рандомайзер

Функция генератора n случайных чисел в диапазоне от 0 до n-1. Но числа могут быть предсказуемы, поскольку генератор сделан на основе времени.

## Реализация (go 1.21.5)

```go
import (
	"math/rand"
	"time"
)

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

## Реализация (go 1.22)

```go
import (
	"math/rand"
)

func RandNumb(n int) <-chan int{
  res := make(chan int)
  go func() {
    for j := 0; j < n; j++ {
      res <- rand.Intn(n)
    }
    close(res)
  }()
  return res
}
```

## Пример
n = 5

На выводе получаем: [2 2 4 1 3]
