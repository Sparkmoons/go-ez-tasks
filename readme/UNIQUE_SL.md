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

## Пример
Слайс А = {1, 2, 3, 4, 5}

Слайс Б = {1, 2, 3, 3, 2, 4, 6, 5, 7}

На выводе получаем: [1 2 3 4 5 6 7]

## Реализация №2 (Для отсортированных слайсов)
```go
func union(A, B []int) []int {
    res := make([]int, 0)

    i, j := 0, 0
    for {
        if A[i] > B[j] {
            res = append(res, B[j])
            j++
        } else if A[i] < B[j] {
            res = append(res, A[i])
            i++
        } else {
            res = append(res, A[i])
            i++
            j++
        }
        if i >= len(A) || j >= len(B) {
            break
        }
    }

    for i < len(A) {
        res = append(res, A[i])
        i++
    } 
    for j < len(B) {
        res = append(res, B[j])
        j++
    }
    return res
}
```