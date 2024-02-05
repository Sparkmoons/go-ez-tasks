## Задача на префикс

Дается слайс из слов, необходимо вывести все слова, которые имеют заданный префикс.

## Реализация №1

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

## Реализация №2

Реализация с использованием функции из пакета strings. 

```go
import "strings"
func PrefCheck(str []string, p string) []string {
  var res []string
  for _, i := range str {
    if strings.HasPrefix(i, p) {
      res = append(res, i)
    }
  }
  return res
}
```