## Реализация LIFO

Дается числовой слайс, необходимо вывести его элементы в обратном порядке. (LIFO - Last Input First Output)

## Реализация №1
Реализация с использованием результирующего слайса. 

```go
func Lifo_1(a []int) []int{
    res := make([]int, len(a), len(a))
    for len(a) > 0 {
        n := len(a) - 1
        res = append(res, a[n])
        a = a[:n]
    }
    return res
}
```

## Реализация №2
Реализация без использования результрующего слайса. 

```go
func Lifo_2(n []int){
    for a, b := 0, len(n)-1; a < b; a, b = a+1, b-1 {
        n[a], n[b] = n[b], n[a]
    }
}
```
