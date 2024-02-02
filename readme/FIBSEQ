## Последовательность чисел Фибоначчи

Возвращает последовательность из n чисел Фибоначчи.

## Реализация

```go
func FibSeq(n int) []int{
  res := make([]int, 0)
  for i := 0; i < n; i++ {
    if i == 0 || i == 1 {
      res = append(res, i)
    }
    a, b := 0, 1
    for n--; n > 0; n-- {
      a, b = b, a + b
      res = append(res, b)
    }
  }
  return res
}
```

## Пример
Пусть n = 8. На выводе будет: [0 1 2 3 5 8 13 21]
