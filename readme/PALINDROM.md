## Проверка слова на палиндром

На вход подается слово. Нужно проверить, является ли оно падиндромом.

## Реализация 

```go
func PalCheck(a string) bool{
  for i := 0; i < len(a) / 2; i++ {
    if a[i] != a[len(i)-1-i] {
      return false
    }
  }
  return true
}
```

## Пример

Слово: topot.  На выводе получаем: true

Слово: topop.  На выводе получаем: false
