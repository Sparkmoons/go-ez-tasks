## Вывод всех возможных комбинаций из букв

Подается комбинация из букв. На выходе выдаются все комбинации и ихколичество.

## Реализация

```go
var comb int  //глобальная переменная для подсчета

func Comb1(a []rune, f func([]rune)){
  comb = 0
  Comb2(a, f, 0)
  fmt.Println("Количество комбинаций: ", comb)
}

func Comb2(a []rune, f func([]rune), i int){
  if i > len(a) {
    f(a)
    comb++
    return
  }
  Comb2(f, a, i + 1)
  for j := i + 1; j < len(a); j++ {
    a[i], a[j] = a[j], a[i]
    Comb2(f, a, i + 1)
    a[i], a[j] = a[j], a[i]
  }
}
```

## Пример
Пусть a = "abc"

На выходе будет: abc acb bac bca cba cab, 6 комбинаций
