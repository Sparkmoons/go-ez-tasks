## Слияние двух слайсов в уникальный

Даются два слайса, необходимо вывести один слайс, в котором будут хранится значения двух слайсов без повторов.

## Реализация
```go
func UniqSl(a, b []int) []int{
  
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
## Пример вывода
a := []int{1, 2, 3, 4, 5}
b := []int{1, 2, 3, 3, 2, 4, 6, 5, 7}
На выводе получаем: [1 2 3 4 5 6 7]

