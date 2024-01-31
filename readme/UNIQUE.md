## Проверка слайса на уникальность компонентов

Дается числовой слайс, необходимо вернуть значение true, если все числа внутри уникальные (без повторов), и false в противном случае.

## Реализация

```go
func UniqCheck(a []int) bool{
  res := make(map[int]bool)
  for _, i := range a {
    _, ok := res[i]
    if ok {
      return false
    }
    res[i] := true
  }
  return true
}
```
