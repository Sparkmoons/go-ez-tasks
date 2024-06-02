## Сортировка

Подается комбинация из букв. На выходе выдаются все комбинации и ихколичество.

## Реализация сортировки пузырьком

```go
func Sort(a []int) []int{
  for i := 0; i < len(a); i++ {
    for j := i + 1; j < len(a); j++ {
      if a[i] > a[j] {
        a[i], a[j] = a[j], a[i]
      }
    }
  }
  return a
}
```

## Реализация сортировки выбором

```go
func Sort(a []int) []int{
  for i := 0; i < len(a); i++ {
    min := i
    for j := i + 1; j < len(a); j++ {
      if a[j] < a[min] {
        min = j
      }
    }
    a[i], a[min] = a[min], a[i]
  }
  return a
}
```

## Реализация сортировки вставкой

```go
func Sort(a []int) []int{
  i := 1
  for i < len(a) {
    j := i
      for j >= 1 && a[i] < a[i-1] {
        a[j], a[j-1] = a[j-1], a[j]
        j--
      }
      i++
  }
  return a
}
```

## Реализация быстрой сортировки

```go
func Sort(a []int) []int {
	if len(a) <= 1 {
		return a
	}

	min := make([]int, 0)
	max := make([]int, 0)
	res := make([]int, 0)
	p := a[0]

	for _, num := range a[1:] {
		if p >= num {
			min = append(min, num)
		} else {
			max = append(max, num)
		}
	}

	sMin := Sort(min)
	sMax := Sort(max)
	res = append(res, sMin...)
	res = append(res, p)
	res = append(res, sMax...)

	//res = append(append(sMin, p), sMax...)

	return res
}
```

## Реализация сортировки через слияние

```go
func Sort1(left, right []int) []int{
  res := make([]int, 0, len(left) + len(right))
  for len(left) > 0 || len(right) > 0 {
    if len(left) == 0 {
      res = append(res, right...)
      break
    } else if len(right) == 0 {
      res = append(res, left...)
      break
    } else if left[0] < right[0] {
      res = append(res, left[0])
      left = left[1:]
    } else if right[0] < left[0] {
      res = append(res, right[0])
      right = right[1:]
    }
  }
  return res
}


func Sort2(arr []int) []int {
  if len(arr) <= 1 {
    return arr
  }
  var left, right []int
  mid := len(arr) / 2

  left = Sort2(arr[:mid])
  right = Sort2(arr[mid:])
  return Sort1(left, right)
}
```

