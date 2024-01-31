package code

import "fmt"

func code(){
  str1 := []string{"arbuz", "ararat", "arka", "zaruba", "paza", "liaar"}
  p := "ar"
  fmt.Printf("Слова с префиксом %s: %s\n", p, PrefCheck(str1, p))
}

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
  
