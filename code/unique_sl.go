package code

import "fmt"

func code(){
  a := []int{1, 2, 3, 4, 5}
	b := []int{1, 2, 3, 3, 2, 4, 6, 5, 7}
  fmt.Printf("Пересечение слайсов: %d\n", UniqSl(a, b))
}

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
