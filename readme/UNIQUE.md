package code

import "fmt"

func code(){
  a := []int{1, 2, 3, 4, 5}
	b := []int{1, 2, 3, 3, 2, 4, 6, 5, 7}
  fmt.Printf("Первый слайс: %t\nВторой слайс: %t\n", UniqCheck(a), UniqCheck(b))
}

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
