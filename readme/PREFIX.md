## Задача на префикс

Дается слайс из слов, необходимо вывести все слова, которые имеют заданный префикс.

## Реализация

```go
func PrefCheck(str []string, p string) []string {
  var res []string
  for _, i := range str {
    match := true
    for j := 0; j < len(p); j++ {
      if i[j] != p[j] {
        match = false
        break
      }
    }
    if match {
      res = append(res, i)
    }
  }
  return res
}
```
