## Разворот слов в строк

На вход подается строка. Нужно перевернуть каждое слово и вывести получившуюся строку. 

## Реализация
```go
import "strings"
func RevWord(str string) string {
    strs := strings.Fields(str)
    reverse := make([]string, 0, len(strs))   
    for _, word := range strs {
        slovo := ""
        for i := len(word)-1; i >= 0; i-- {
            slovo = string(word[i])
        }
        reverse = append(reverse, slovo)
    }
    return strings.Join(reverse, " ")
}
```

## Пример
Входная строка: "Another one on Go"

На выводе получаем: "rehtonA eno no oG"
