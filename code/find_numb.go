package code

import "fmt"

func code(){
  c := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
  pp := 8
  fmt.Printf("Индекс числа %d: %d\n", pp, FindNumb(c, pp))
}

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
