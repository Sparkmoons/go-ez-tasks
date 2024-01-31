## Слияние двух слайсов в уникальный

Даются два слайса, необходимо вывести один слайс, в котором будут хранится значения двух слайсов без повторов.

## Реализация
```go
func Cross(a, b []int) []int{
  
  res := make([]int, 0)
  counter := make(map[int]int)
  
  for _, i := range a {
    if _, ok := counter[i]; !ok {
      counter[i] = 1
      res = append(res, i)
    }
  }

  for _, j := range b {
    if c, ok := counter[j]; !ok {
      counter[j] -= 1
      res = append(res, j)
    }
  }
  return res  
}
```
