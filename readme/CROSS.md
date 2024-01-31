## Вывод пересечения двух слайсов

Даются два слайса, необходимо вывести один слайс, в котором будут хранится общие значения для первого и второго слайса.

## Реализация
```go
func Cross(a, b []int) []int{
  
  res := make([]int, 0)
  counter := make(map[int]int)
  
  for _, i := range a {
    if _, ok := counter[i]; !ok {
      counter[i] = 1
    } else {
      counter[i] += 1
    }
  }

  for _, j := range b {
    if c, ok := counter[j]; c > 0 && ok {
      res = append(res, j)
      counter[j] -= 1
    }
  }
  return res  
}
```
