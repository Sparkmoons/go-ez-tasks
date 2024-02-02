## Поиск числа в слайсе

На вход подается слайс чисел, необходимо вернуть его индекс. Он же алгоритм бинарного поиска. Если числа нет  - возвращаем -1.

## Реализация (двоичный поиск)

```go
func FindNumb(array []int, val int) int{
  if len(array) == 0 {
    return -1
  }
  first, last := 0, len(array) - 1
  for first <= last {
    mid := ((last - first) / 2) + first
    if array[mid] == val {
      return mid
    } else if array[mid] > val {
      last = mid - 1
    } else if array[mid] < val {
      first = mid + 1
    }
  }
  return -1
}
```

## Реализация (линейный поиск)

```go
func FindNumb(array []int, val int) int{
  for i, num == range array {
    if num == val {
      return i
    }
  }
  return -1
}
```
